package uploader

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	hakka_config "github.com/rice9547/hakka_story/config"
)

type Client struct {
	cli  *s3.Client
	conf hakka_config.SpaceConfig
}

func New(cfg hakka_config.SpaceConfig, cli *s3.Client) *Client {
	return &Client{
		conf: cfg,
		cli:  cli,
	}
}

func NewS3Client(cfg hakka_config.SpaceConfig) (*s3.Client, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKey,
			cfg.SecretKey,
			"",
		)),
		config.WithRegion(cfg.Region),
	)
	if err != nil {
		return nil, fmt.Errorf("create AWS config failed: %v", err)
	}

	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.Endpoint)
		o.UsePathStyle = true
	})

	return s3Client, nil
}

func (c *Client) UploadImage(ctx context.Context, file io.Reader, filename, contentType string) (string, error) {
	return c.upload(ctx, file, c.conf.ImageBucket, filename, contentType)
}

func (c *Client) UploadAudio(ctx context.Context, file io.Reader, filename, contentType string) (string, error) {
	return c.upload(ctx, file, c.conf.AudioBucket, filename, contentType)
}

func (c *Client) upload(ctx context.Context, file io.Reader, bucket, filename, contentType string) (string, error) {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	}

	_, err := c.cli.PutObject(ctx, input)
	if err != nil {
		return "", fmt.Errorf("upload image failed: %v", err)
	}

	url := fmt.Sprintf("%s/%s/%s", c.conf.Endpoint, bucket, filename)
	return url, nil
}
