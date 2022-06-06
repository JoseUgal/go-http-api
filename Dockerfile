FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/JoseUgal/go-http-api
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/go-http-api cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/go-http-api /go/bin/go-http-api
ENTRYPOINT ["/go/bin/go-http-api"]