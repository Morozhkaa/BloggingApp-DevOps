FROM golang:1.19-buster as builder

WORKDIR /auth-service
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o application

FROM alpine:3.15.4
WORKDIR /auth-service
COPY --from=builder /auth-service/application /auth-service/application
COPY *.yaml ./
CMD ["/auth-service/application"]
