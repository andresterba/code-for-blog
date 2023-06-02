FROM golang:1.20-alpine AS builder

COPY . /app
WORKDIR /app

RUN go get ./...

RUN go build -o test ./...



FROM scratch

COPY --from=builder /app/test /test

ENTRYPOINT ["/test"]