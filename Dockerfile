FROM golang:1.21.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM ubuntu:latest

COPY --from=builder /app/build/main /main

EXPOSE 8000

CMD ["/main"]