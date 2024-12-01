FROM golang:1.23.2-alpine3.19 as builder

WORKDIR /build

RUN apk --no-cache add make

COPY . .

RUN go mod download

ARG RELEASE
ARG GIT_HASH
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Release=${RELEASE} -X main.GitHash=${GIT_HASH}" -o /build/bin/service ./cmd/service

FROM alpine:3.19

RUN mkdir -p /app
COPY --from=builder /build/bin/service /app/service

ENTRYPOINT ["/app/service"]
