#!/bin/sh

# compiling the app
go build -o build/versioncompare
chmod +x build/versioncompare

# running the app
# Default cron schedule : every minute
APP_CRON_SCHEDULE="${1:-0 */1 * * * *}" \
TEST_URL="${2:-https://horizon-testnet.stellar.org/}" \
PROD_URL="${3:-https://horizon.stellar.org/}" \
./build/versioncompare