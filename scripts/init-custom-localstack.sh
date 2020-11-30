#!/bin/sh

set -eux

DIR=$(cd `dirname $0`/../; pwd)
cd $DIR

echo '--- create s3 and put lambda key ---'
awslocal s3api create-bucket --bucket lambda-bucket
./scripts/upload-lambda-to-s3-localstack.sh

echo '--- create s3 and put sample custom excel file ---'
awslocal s3api create-bucket --bucket custom-bucket
awslocal s3 cp test/data/Book1.xlsx s3://custom-bucket

echo '--- validate template ---'
awslocal cloudformation validate-template \
  --template-body file://template.yaml

echo '--- create localstack ---'
awslocal cloudformation create-stack \
  --stack-name custom \
  --template-body file://template.yaml \
  --capabilities CAPABILITY_IAM

awslocal cloudformation wait stack-create-complete --stack-name custom
