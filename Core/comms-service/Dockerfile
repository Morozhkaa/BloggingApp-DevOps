FROM golang:1.19-buster as builder

WORKDIR /comm-service
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o application

FROM alpine:3.15.4
WORKDIR /comm-service
COPY --from=builder /comm-service/application /comm-service/application
COPY *.yaml ./
CMD ["/comm-service/application"]
