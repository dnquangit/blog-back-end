package uploadprovider

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go-module/component"
)

type s3Provider struct {
	Bucket  string
	Key     string
	Secret  string
	Region  string
	Url     string
	Session *session.Session
}

func NewS3Provider(bucket string, key string, secret string, region string, url string) (*s3Provider, error) {
	s3Session, err := session.NewSession(&aws.Config{
		Region: &region,
		Credentials: credentials.NewStaticCredentials(
			key,    // Access key ID
			secret, // Secret access key
			""),    // Token can be ignore
	})

	if err != nil {
		return nil, err
	}

	return &s3Provider{
		Bucket:  bucket,
		Key:     key,
		Secret:  secret,
		Region:  region,
		Session: s3Session,
		Url:     url,
	}, nil
}

func (provider *s3Provider) UploadFile(ctx context.Context, fileName string, fileData []byte) (string, error) {
	if provider.Session == nil {
		return "", component.NewAppError("Cannot connect to s3 provider", component.ErrInternal.String(), "The s3 Session is nil")
	}

	uploader := s3manager.NewUploader(provider.Session)
	reader := bytes.NewReader(fileData)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(provider.Bucket),
		Key:    aws.String(fileName),
		Body:   reader,
	})

	if err != nil {
		return "", component.NewAppError("Cannot connect to s3 provider", component.ErrInternal.String(), err.Error())
	}

	url := "https://" + provider.Url + "/" + fileName
	return url, nil
}
