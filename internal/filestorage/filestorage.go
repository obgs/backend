package filestorage

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	miniocreds "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

type FileStorageService struct {
	s3Client        *s3.Client
	s3PresignClient *s3.PresignClient
	bucket          string
	usingMinio      bool
	minioClient     *minio.Client
}

const BUCKET_CORS_MAX_AGE_SECONDS = 3000

// NewFileStorageService creates a new FileStorageService
func NewFileStorageService(accessKeyID,
	secretAccessKey,
	region,
	endpoint,
	bucket string,
	usingMinio bool,
) (*FileStorageService, error) {
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""),
		Region:      region,
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           endpoint,
					SigningRegion: region,
				}, nil
			}),
	}
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if usingMinio {
			o.UsePathStyle = true
		}
	})
	s3PresignClient := s3.NewPresignClient(s3Client)
	minioEndpoint := endpoint
	minioEndpoint = strings.TrimPrefix(strings.TrimPrefix(minioEndpoint, "http://"), "https://")
	minioClient, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  miniocreds.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &FileStorageService{
		s3Client,
		s3PresignClient,
		bucket,
		usingMinio,
		minioClient,
	}, nil
}

// SignUploadURL signs a URL for uploading a file to S3
func (s *FileStorageService) SignUploadURL(ctx context.Context) (*model.UploadURL, error) {
	key := uuid.New()
	req, err := s.s3PresignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key.String()),
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return nil, err
	}

	headers := make([]*model.Header, 0, len(req.SignedHeader))
	for k, v := range req.SignedHeader {
		for _, vv := range v {
			headers = append(headers, &model.Header{
				Key:   k,
				Value: vv,
			})
		}
	}

	return &model.UploadURL{
		URL:     req.URL,
		Headers: headers,
	}, nil
}

func (s *FileStorageService) CreateBucket(ctx context.Context) error {
	resp, err := s.s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return err
	}
	for _, bucket := range resp.Buckets {
		if *bucket.Name == s.bucket {
			return nil
		}
	}
	if s.usingMinio {
		err = s.minioClient.MakeBucket(ctx, s.bucket, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
		err = s.minioClient.SetBucketPolicy(
			ctx,
			s.bucket,
			//nolint:lll
			`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetBucketLocation","s3:ListBucket"],"Resource":["arn:aws:s3:::`+s.bucket+`"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::`+s.bucket+`/*"]}]} `,
		)
	} else {
		_, err = s.s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(s.bucket),
			ACL:    types.BucketCannedACLPublicRead,
		})

		if err != nil {
			return err
		}
		_, err = s.s3Client.PutBucketCors(ctx, &s3.PutBucketCorsInput{
			Bucket: aws.String(s.bucket),
			CORSConfiguration: &types.CORSConfiguration{
				CORSRules: []types.CORSRule{
					{
						AllowedHeaders: []string{"*"},
						AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
						AllowedOrigins: []string{"*"},
						ExposeHeaders:  []string{"ETag"},
						MaxAgeSeconds:  BUCKET_CORS_MAX_AGE_SECONDS,
					},
				},
			},
		})
	}

	return err
}
