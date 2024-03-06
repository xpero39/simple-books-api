FROM golang:1.22.1-alpine as builder
COPY go.mod go.sum /go/src/github.com/xpero39/simple-books-api/
WORKDIR /go/src/github.com/xpero39/simple-books-api
RUN go mod download
COPY . /go/src/github.com/xpero39/simple-books-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/simple-books-api github.com/xpero39/simple-books-api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/xpero39/simple-books-api/build/simple-books-api /usr/bin/simple-books-api
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/simple-books-api"]