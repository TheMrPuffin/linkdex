FROM golang:1.20.2-alpine3.17 AS builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o api .

FROM alpine:latest

WORKDIR /

COPY --from=builder /build/api .

EXPOSE 8080

CMD [ "/api" ]