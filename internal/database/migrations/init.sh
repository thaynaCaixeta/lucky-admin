#!/bin/bash
set -euo pipefail

export AWS_ACCESS_KEY_ID=fake
export AWS_SECRET_ACCESS_KEY=fake

ENDPOINT="http://localhost:8000"
REGION="us-west-2"
TABLE_NAME="GameSystem"

echo "Running inside container? Hostname: $(hostname)"

echo "Waiting for DynamoDB Local to be ready..."
until curl -s "$ENDPOINT" > /dev/null; do
  echo "Waiting for DynamoDB Local to be ready..."
  sleep 1
done

# Create table
echo "Creating table '$TABLE_NAME'..."
if ! aws dynamodb create-table \
  --cli-input-json file://internal/database/migrations/01_create_table.json \
  --endpoint-url "$ENDPOINT" \
  --region "$REGION"; then
  echo "❌ Failed to create table '$TABLE_NAME'. Check JSON and DynamoDB logs."
  exit 1
fi

# Seed default admin
echo "Adding default admin..."
jq -c '.[]' internal/database/migrations/02_seed_default_admin.json | while read -r item; do
  echo "{\"TableName\": \"$TABLE_NAME\", \"Item\": $item}" > internal/database/migrations/tmp-put-item.json

  aws dynamodb put-item \
    --cli-input-json file://internal/database/migrations/tmp-put-item.json \
    --endpoint-url "$ENDPOINT" \
    --region "$REGION"

  echo "Inserted: $item"
done
rm -f internal/database/migrations/tmp-put-item.json

# Seed test data
echo "Seeding test data..."
jq -c '.[]' internal/database/migrations/03_seed_test_data.json | while read -r row; do
  echo "{\"TableName\": \"$TABLE_NAME\", \"Item\": $row}" > internal/database/migrations/tmp-put-item.json

  aws dynamodb put-item \
    --cli-input-json file://internal/database/migrations/tmp-put-item.json \
    --endpoint-url "$ENDPOINT" \
    --region "$REGION"

  echo "Inserted: $row"
done
rm -f internal/database/migrations/tmp-put-item.json

echo "✅ All seed data loaded!"