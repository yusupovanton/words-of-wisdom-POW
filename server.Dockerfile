FROM golang:1.23.2 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server ./cmd/server/main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/server /app/server
RUN chmod +x /app/server

EXPOSE 8080
CMD ["/app/server"]
