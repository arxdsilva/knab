FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN go build -o knab ./cmd/api/main.go
EXPOSE 8888
ENTRYPOINT ./knab