module github.com/dinhnguyen7025/example-using-sqs-with-localstack

go 1.15

replace github.com/dinhnguyen7025/example-using-sqs-with-localstack/ => ./

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.1
	github.com/aws/aws-lambda-go v1.19.1
	github.com/aws/aws-sdk-go v1.35.10
)
