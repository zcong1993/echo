FROM zcong/golang:1.10.3 AS build
WORKDIR /go/src/github.com/zcong1993/echo
COPY . .
RUN dep ensure -vendor-only -v && \
    CGO_ENABLED=0 go build -o ./bin/echo main.go

FROM alpine:3.7
WORKDIR /opt
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/zcong1993/echo/bin/* /usr/bin/
EXPOSE 8080
CMD ["echo"]
