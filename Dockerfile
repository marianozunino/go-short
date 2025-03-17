FROM golang:1.24.1-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go run github.com/a-h/templ/cmd/templ generate
RUN go run github.com/sqlc-dev/sqlc/cmd/sqlc generate
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o /go/bin/go-short ./cmd

# Install migrate CLI
RUN go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:3.19
ARG PUID=1000
ARG PGID=1000
ARG PORT=1323

RUN apk add --no-cache sqlite-libs ca-certificates

COPY --from=builder /go/bin/go-short /app/go-short
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY migrations /app/migrations
COPY docker-entrypoint.sh /app/

RUN mkdir -p /data && \
    chown -R 1000:1000 /data /app && \
    chmod +x /app/docker-entrypoint.sh

VOLUME ["/data"]
EXPOSE ${PORT}

ENV BASE_DOMAIN="http://localhost:1323" \
    PORT="1323" \
    DB_PATH="/data/go-short.sqlite"

USER ${PUID}:${PGID}
WORKDIR /app
ENTRYPOINT ["/app/docker-entrypoint.sh"]
