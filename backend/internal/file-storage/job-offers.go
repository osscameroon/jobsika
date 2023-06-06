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
func (fs FileStorage) UploadJobOfferCompanyPicture(file []byte, jobOfferID int64) (string, error) {
	fileName := fmt.Sprintf("%s/%d/%d.png", JobOffersCompanyPicturesFolder, jobOfferID, jobOfferID)
	_, err := fs.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(fs.config.Bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		return "", err
	}

	return fileName, nil
}

// DownloadJobOfferCompanyPicture downloads a job offer company picture from the file storage.
func (fs FileStorage) DownloadJobOfferCompanyPicture(jobOfferID int64) (*s3.GetObjectOutput, error) {
	fileName := fmt.Sprintf("%s/%d/%d.png", JobOffersCompanyPicturesFolder, jobOfferID, jobOfferID)
	result, err := fs.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(fs.config.Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
