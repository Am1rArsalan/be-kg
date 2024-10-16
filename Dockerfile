FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-server ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-server .

EXPOSE 8080

CMD ["./go-server"]

