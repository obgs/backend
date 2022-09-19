package filestorage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type FileStorageService struct {
	client *s3.PresignClient
	bucket string
}

// NewFileStorageService creates a new FileStorageService
func NewFileStorageService(accessKeyID, secretAccessKey, region, endpoint, bucket string, usingMinio bool) *FileStorageService {
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""),
		Region:      region,
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           endpoint,
				SigningRegion: region,
			}, nil
		}),
	}
	client := s3.NewPresignClient(s3.NewFromConfig(cfg, func(o *s3.Options) {
		if usingMinio {
			o.UsePathStyle = true
		}
	}))
	return &FileStorageService{
		client,
		bucket,
	}
}

// SignUploadURL signs a URL for uploading a file to S3
func (s *FileStorageService) SignUploadURL(ctx context.Context) (string, error) {
	key := uuid.New()
	req, err := s.client.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key.String()),
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}
