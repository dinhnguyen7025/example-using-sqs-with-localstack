package s3_downloader

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	tempArtifactPath = "/tmp/artifact/"
	dirPerm          = 0755
)

type Downloader struct {
	s3manager   s3manager.Downloader
	bucket, key string
}

// NewDownloader func
func NewDownloader(s *session.Session, bucket, key string) *Downloader {
	return &Downloader{
		s3manager: *s3manager.NewDownloader(s),
		bucket:    bucket,
		key:       key,
	}
}

// Download func
func (d *Downloader) Download(downloadDir string) (string, error) {
	file, err := os.Create(downloadDir + d.key)
	if err != nil {
		return "", err
	}
	defer file.Close()

	numBytes, err := d.s3manager.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(d.bucket),
			Key:    aws.String(d.key),
		})

	if err != nil {
		return "", err
	}

	log.Println("Downloaded", file.Name(), numBytes, "bytes")

	return file.Name(), nil
}

// PrepareTmpDownloadDir func
func (d *Downloader) PrepareTmpDownloadDir() (string, error) {
	now := strconv.Itoa(int(time.Now().UnixNano()))
	tmpDir := tempArtifactPath + now + "/"
	if _, err := os.Stat(tmpDir); err == nil {
		if err := os.RemoveAll(tmpDir); err != nil {
			return "", err
		}
	}

	if err := os.MkdirAll(tmpDir, dirPerm); err != nil {
		return "", err
	}

	return tmpDir, nil
}
