package filestorage

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	JobOffersCompanyPicturesFolder = "/job-offers"
)

// UploadJobOfferCompanyPicture uploads a job offer company picture to the file storage.
func (fs FileStorage) UploadJobOfferCompanyPicture(file []byte, jobOfferID int64, mimeType string, extension string) (string, error) {
	// get the right file extension based on content type of the file

	fileName := fmt.Sprintf("%s/%d/%d%s", JobOffersCompanyPicturesFolder, jobOfferID, jobOfferID, extension)
	_, err := fs.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(fs.config.Bucket),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(file),
		ContentType: aws.String(mimeType),
	})
	if err != nil {
		return "", err
	}

	return fileName, nil
}

// DownloadJobOfferCompanyPicture downloads a job offer company picture from the file storage.
func (fs FileStorage) DownloadJobOfferCompanyPicture(location string) (*s3.GetObjectOutput, error) {
	result, err := fs.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(fs.config.Bucket),
		Key:    aws.String(location),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
