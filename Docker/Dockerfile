FROM golang:1.21-alpine AS builder

COPY . .
RUN apk update && \
    go build -o /target/whatismyip main.go

FROM alpine

RUN apk update && \
    apk add --no-cache bash
WORKDIR /home/admin
COPY --from=builder /target/ .
EXPOSE 8080
CMD ["/home/admin/whatismyip"]
