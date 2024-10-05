FROM golang:1.22 AS builder

RUN mkdir /build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o smolurl

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /build /app

CMD ["./smolurl"]
