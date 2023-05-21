FROM golang:1.19-buster as builder

WORKDIR /post-service
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o application

FROM alpine:3.15.4
WORKDIR /post-service
COPY --from=builder /post-service/application /post-service/application
COPY *.yaml ./
CMD ["/post-service/application"]
