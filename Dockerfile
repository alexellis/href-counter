FROM alpine:3.10
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY vendor vendor
COPY app    .

CMD ["./app"]
