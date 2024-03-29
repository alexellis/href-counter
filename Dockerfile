FROM alpine:3.17.1
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY app    .

CMD ["./app"]
