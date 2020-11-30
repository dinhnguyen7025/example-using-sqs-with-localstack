package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	custom_hander "github.com/dinhnguyen7025/example-using-sqs-with-localstack/internal/pkg/custom"
)

func main() {
	lambda.Start(custom_hander.NewCustomHandler().Process)
}
