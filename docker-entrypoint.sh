#!/bin/sh
set -e

# Run migrations using golang-migrate
echo "Running database migrations..."
migrate -path /app/migrations -database "sqlite3://${DB_PATH}" up

# Start the application
echo "Starting application..."
exec /app/go-short
