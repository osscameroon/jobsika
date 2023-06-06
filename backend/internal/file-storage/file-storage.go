package filestorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//
type FileStorage struct {
	//The database client
	s3Client *s3.S3
}

type FileStorageOptions struct {
	S3AccessKey string
	S3SecretKey string
	Endpoint    string
}

//NewFileStorage returns a new FileStorage client
func NewFileStorage(opt FileStorageOptions) (*FileStorage, error) {
	config := &aws.Config{
		Endpoint:         aws.String(opt.Endpoint),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(opt.S3AccessKey, opt.S3SecretKey, ""),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess)

	return &FileStorage{
		s3Client: s3Client,
	}, nil
}
