package database

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	appconfig "github.com/thaynaCaixeta/lucky-admin/internal/config"
)

func NewLocalDynamoClient(ctx context.Context, cfg appconfig.DynamoDBConfig) (*dynamodb.Client, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.LocalRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("fakeId", "fakeSecret", "")),
		config.WithEndpointResolver(localDynamoEndpointResolver(cfg.LocalEndpoint, cfg.LocalRegion)),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}
	return dynamodb.NewFromConfig(awsCfg), nil
}

func localDynamoEndpointResolver(endpoint, region string) aws.EndpointResolver {
	return aws.EndpointResolverFunc(func(service, _ string) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID {
			return aws.Endpoint{
				URL:           endpoint,
				SigningRegion: region,
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown service: %s", service)
	})
}
