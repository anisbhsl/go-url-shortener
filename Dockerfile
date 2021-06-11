FROM alpine:latest as base

FROM golang:1.14.1-buster AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux  go build -a -installsuffix cgo -o server server.go

FROM base as prod
WORKDIR /app
COPY --from=builder /app/server .
ENTRYPOINT ["./server"]