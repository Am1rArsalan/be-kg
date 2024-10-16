FROM golang:1.23-alpine as builder

WORKDIR /app

# Copy only go.mod first (without go.sum) to avoid errors when go.sum is missing
COPY go.mod ./

# Download dependencies (this will create go.sum if it's missing)
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o go-server ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-server .

EXPOSE 8080

CMD ["./go-server"]

