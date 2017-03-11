package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"net/url"

	"strings"

	"encoding/json"

	"golang.org/x/net/html"
)

type scrapeDataStore struct {
	Internal int `json:"internal"`
	External int `json:"external"`
}

func main() {
	urlIn := os.Getenv("url")
	if len(urlIn) == 0 {
		log.Fatalln("Need a valid url as an env-var.")
	}

	siteUrl, parseErr := url.Parse(urlIn)

	if parseErr != nil {
		log.Fatalln(parseErr)
	}

	resp, err := http.Get(urlIn)
	if err != nil {
		log.Fatalln(err)
	}

	scrapeData := &scrapeDataStore{}

	tokenizer := html.NewTokenizer(resp.Body)
	end := false
	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.StartTagToken:
			// fmt.Println(tt)
			token := tokenizer.Token()
			switch token.Data {
			case "a":

				for _, attr := range token.Attr {

					if attr.Key == "href" {
						link := attr.Val

						parsedLink, parseLinkErr := url.Parse(link)

						// if strings.Index(link, "#") == 0 {
						// 	fmt.Println(link)
						// } else {
						// 	fmt.Println(parsedLink)
						// }
						if parseLinkErr == nil {
							if parsedLink.Host == siteUrl.Host || strings.Index(link, "#") == 0 || len(parsedLink.Host) == 0 {
								scrapeData.Internal++
							} else {
								scrapeData.External++
							}
						}

						if parseLinkErr != nil {
							fmt.Println("Can't parse: " + token.Data)
						}
					}
				}
				break
			}
		case tt == html.ErrorToken:
			end = true
			break
		}
		if end {
			break
		}
	}
	data, _ := json.Marshal(&scrapeData)
	fmt.Println(string(data))
}
