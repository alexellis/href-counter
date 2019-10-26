# href-counter

> A Golang example application which counts internal vs. external hrefs within a page to rate SEO.

The `golang.org/x/net/html` package is used to iterate through all the HTML tokens in the web-page. It provides a working example of parsing HTML piece-by-piece. 

This can be built with the Dockerfile in the repository or through `go get`.

### References

Used in these two blog posts:

* [One-shot containers on Docker Swarm](http://blog.alexellis.io/containers-on-swarm/)
* [Builder pattern vs. Multi-stage builds in Docker](http://blog.alexellis.io/mutli-stage-docker-builds/)

### Running the example

* With Go

    ```sh
    go build

    $ url=http://blog.alexellis.io/ ./href-counter
    {"internal":40,"external":2}

    $ url=http://blog.alexellis.io/golang-json-api-client/  ./href-counter
    {"internal":17,"external":15}
    ```

* Build with multi-stage build

    ```sh
    ./build.sh
    ```

* Run with Docker

    ```sh
    docker run -e url=https://www.alexellis.io/ -ti alexellis2/href-counter:0.1.0
    {"internal":6,"external":11}
    ```

* Build with multiple Dockerfiles

    ```sh
    ./build-multi-dockerfiles.sh
    ```
