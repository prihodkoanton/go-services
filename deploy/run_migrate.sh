#!/bin/sh
set -e

echo "=== ENV ==="
env | sort

echo "=== LS /migrations ==="
ls -lah /migrations

echo "Running migrate up..."

migrate \
  -verbose \
  -path=/migrations \
  -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$DB_HOST:$DB_PORT/$POSTGRES_DB?sslmode=$DB_SSLMODE" \
  up || true

echo "Migrations finished"