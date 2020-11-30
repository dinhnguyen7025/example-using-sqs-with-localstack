package custom

import (
	"context"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	s3_downloader "github.com/dinhnguyen7025/example-using-sqs-with-localstack/pkg/s3-downloader"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Handler struct {
	Queue        *sqs.SQS
	S3Downloader *s3_downloader.Downloader
}

const (
	FileName = "Book1.xlsx"
)

//NewCustomHandler func
func NewCustomHandler() *Handler {
	endpoint := os.Getenv("LOCALSTACK_ENDPOINT")
	customBucket := os.Getenv("CUSTOM_BUCKET")

	if endpoint != "" && customBucket != "" {
		// Create S3 service client
		s3CustomResolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
			if service == endpoints.S3ServiceID {
				return endpoints.ResolvedEndpoint{
					URL: endpoint,
				}, nil
			}

			return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
		}

		sess := session.Must(session.NewSession(&aws.Config{
			EndpointResolver: endpoints.ResolverFunc(s3CustomResolver),
			S3ForcePathStyle: aws.Bool(true), //https://github.com/aws/aws-sdk-go/issues/2743
		}))

		return &Handler{
			Queue:        sqs.New(sess, &aws.Config{}),
			S3Downloader: s3_downloader.NewDownloader(sess, customBucket, FileName),
		}

	}

	return &Handler{
		Queue:        sqs.New(session.New(), &aws.Config{}),
		S3Downloader: s3_downloader.NewDownloader(session.New(), customBucket, FileName),
	}
}

//GetMessages func
func (customHandler *Handler) GetMessages() (*sqs.ReceiveMessageOutput, error) {
	msgResult, err := customHandler.Queue.ReceiveMessage(&sqs.ReceiveMessageInput{})

	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

//ReadCustomFile func
func (customHandler *Handler) ReadCustomFileSheetName(filePath string, fileName string) (string, error) {

	file, err := excelize.OpenFile(filePath + fileName)
	if err != nil {
		log.Println(err)
		return "", err
	}

	firstSheetName := file.WorkBook.Sheets.Sheet[0].Name

	return firstSheetName, nil
}

//Process func
func (customHandler *Handler) Process(ctx context.Context, sqsEvent events.SQSEvent) error {
	// Read file from s3
	// Create tmpDir
	tmpDir, err := customHandler.S3Downloader.PrepareTmpDownloadDir()
	if err != nil {
		log.Fatal(err)
	}

	// Downloaded file from s3
	_, err = customHandler.S3Downloader.Download(tmpDir)
	if err != nil {
		log.Fatal(err)
	}

	sheetName, err := customHandler.ReadCustomFileSheetName(tmpDir, FileName)
	log.Println("download file successfully with sheet name is " + sheetName)

	return nil
}
