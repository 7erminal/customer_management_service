FROM golang:1.21-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/customer_management_service .

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata \
	&& addgroup -S appgroup \
	&& adduser -S appuser -G appgroup \
	&& mkdir -p /logs /app/conf /app/swagger /app/uploads \
	&& chown -R appuser:appgroup /logs /app

COPY --from=builder /out/customer_management_service /app/customer_management_service
COPY conf /app/conf
COPY swagger /app/swagger

COPY <<'EOF' /app/entrypoint.sh
#!/bin/sh
set -eu

if [ -n "${APP_HTTP_PORT:-}" ]; then
  export BEEGO_HTTPPORT="${APP_HTTP_PORT}"
fi

if [ -n "${BEEGO_RUNMODE:-}" ]; then
  export BEEGO_RUNMODE
fi

exec /app/customer_management_service
EOF

RUN chmod +x /app/entrypoint.sh

USER appuser

EXPOSE 5083

ENTRYPOINT ["/app/entrypoint.sh"]