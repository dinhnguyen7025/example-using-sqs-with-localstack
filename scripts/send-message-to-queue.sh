#!/bin/sh

set -eux

echo '--- update newest lambda function code ---'
./scripts/build-local.sh
./scripts/update-lambda-func-code.sh

echo '--- push message to queue ---'
awslocal sqs send-message \
  --queue-url http://localhost:4566/000000000000/example-queue \
  --message-body "test message"

