package main

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Presigner interface {
	PresignGetObject(context.Context, *s3.GetObjectInput, ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignPutObject(context.Context, *s3.PutObjectInput, ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

// S3Presigner encapsulates the Amazon Simple Storage Service (Amazon S3) presign actions
// used in the examples.
// It contains PresignClient, a client that is used to presign requests to Amazon S3.
// Presigned requests contain temporary credentials and can be made from any HTTP client.
type S3Presigner struct {
	client *s3.PresignClient
	region string
}

func NewS3Presigner(client *s3.PresignClient, region string) S3Presigner {
	return S3Presigner{client, region}
}

// GetObject makes a presigned request that can be used to get an object from a bucket.
// The presigned request is valid for the specified number of seconds.
func (p S3Presigner) GetObject(ctx context.Context, bucket, key string, exp time.Duration) (*v4.PresignedHTTPRequest, error) {
	input := &s3.GetObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)}

	return p.client.PresignGetObject(ctx, input, s3.WithPresignExpires(exp), withRegion(p.region))
}

// PutObject makes a presigned request that can be used to put an object in a bucket.
// The presigned request is valid for the specified number of seconds.
func (p S3Presigner) PutObject(ctx context.Context, bucket, key string, exp time.Duration) (*v4.PresignedHTTPRequest, error) {
	input := &s3.PutObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)}

	return p.client.PresignPutObject(ctx, input, s3.WithPresignExpires(exp), withRegion(p.region))
}

// withRegion is a helper function that adds the region to the PresignOptions.
func withRegion(region string) func(*s3.PresignOptions) {
	return func(o *s3.PresignOptions) {
		o.ClientOptions = append(o.ClientOptions,
			func(o *s3.Options) { o.Region = region },
		)
	}
}
