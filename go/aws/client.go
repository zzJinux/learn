package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/cockroachdb/errors"
)

// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/endpoints/#immutable-endpoint

type staticResolver struct{}

func (*staticResolver) ResolveEndpoint(_ context.Context, params s3.EndpointParameters) (
	smithyendpoints.Endpoint, error,
) {
	if params.Endpoint == nil {
		return smithyendpoints.Endpoint{}, fmt.Errorf("endpoint is nil")
	}
	u, err := url.Parse(*params.Endpoint)
	if err != nil {
		return smithyendpoints.Endpoint{}, err
	}
	if params.Bucket != nil {
		u = u.JoinPath(*params.Bucket)
	}
	return smithyendpoints.Endpoint{URI: *u}, nil
}

func NewS3Client(profile string) (*s3.Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return nil, err
	}
	client := s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.EndpointResolverV2 = &staticResolver{}
	})
	return client, nil
}

func CopyObjectClientSide(ctx context.Context, dstClient *s3.Client, dstBucket, dstKey string, srcClient *s3.Client, srcBucket, srcKey string) error {
	get, err := srcClient.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(srcBucket),
		Key:    aws.String(srcKey),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	defer get.Body.Close()

	_, err = manager.NewUploader(dstClient, func(u *manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024
	}).Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(dstBucket),
		Key:    aws.String(dstKey),
		Body:   get.Body,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
