package database

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	appconfig "github.com/thaynaCaixeta/lucky-admin/internal/config"
)

func NewProdDynamoClient(ctx context.Context, cfg appconfig.DynamoDBConfig) (*dynamodb.Client, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.AwsRegion),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}
	return dynamodb.NewFromConfig(awsCfg), nil
}
