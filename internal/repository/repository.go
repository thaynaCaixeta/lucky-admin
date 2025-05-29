package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/thaynaCaixeta/lucky-admin/internal/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbTypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Repository interface {
	SaveGame(ctx context.Context, numRounds int, closesAt time.Time, createdBy string) (*domain.Game, error)
}

type repo struct {
	cli *dynamodb.Client
}

func NewRepository(cli *dynamodb.Client) Repository {
	return &repo{
		cli: cli,
	}
}

type GameItem struct {
	PK               string `dynamodbav:"PK"`
	SK               string `dynamodbav:"SK"`
	Id               string `dynamodbav:"id"`
	NumRounds        int    `dynamodbav:"num_rounds"`
	CreatedAt        string `dynamodbav:"created_at"`
	ClosesAt         string `dynamodbav:"closes_at"`
	CompletionStatus string `dynamodbav:"completion_status"`
	CreatedBy        string `dynamodbav:"created_by"`
}

type adminItem struct {
	Id        string `dynamodbav:"id"`
	Username  string `dynamodbav:"username"`
	Password  string `dynamodbav:"pass"`
	CreatedAt string `dynamodbav:"created_at"`
	IsActive  bool   `dynamodbav:"is_active"`
}

func (r *repo) SaveGame(
	ctx context.Context,
	numRounds int,
	closesAt time.Time,
	createdBy string,
) (*domain.Game, error) {
	resp, err := r.cli.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String("GameSystem"),
		IndexName:              aws.String("UsernameIndex"),
		KeyConditionExpression: aws.String("username = :u"),
		ExpressionAttributeValues: map[string]dynamodbTypes.AttributeValue{
			":u": &dynamodbTypes.AttributeValueMemberS{Value: createdBy},
		},
		Limit: aws.Int32(1),
	})
	if err != nil {
		return nil, NewDatabaseError("error while retrieving the admin from the database", err)
	}
	if len(resp.Items) == 0 {
		return nil, NewAdminNotFoundError(createdBy)
	}
	// Unmarshal response
	var admItem adminItem
	if err = attributevalue.UnmarshalMap(resp.Items[0], &admItem); err != nil {
		return nil, NewDatabaseError("failed to parse admin", err)
	}

	createdAtAdm, err := time.Parse(time.RFC3339, admItem.CreatedAt)
	if err != nil {
		return nil, NewDatabaseError("invalid admin timestamp", err)
	}
	if !admItem.IsActive {
		return nil, NewInvalidAdminError("admin is not active", err)
	}

	// Generate a new game and save
	gameId := uuid.New().String()
	now := time.Now().UTC()

	gameItem := GameItem{
		PK:               "GAME#" + gameId,
		SK:               "METADATA",
		Id:               gameId,
		NumRounds:        numRounds,
		CreatedAt:        now.Format(time.RFC3339),
		ClosesAt:         closesAt.Format(time.RFC3339),
		CompletionStatus: domain.OnGoing.String(),
		CreatedBy:        admItem.Username,
	}

	item, err := attributevalue.MarshalMap(gameItem)
	if err != nil {
		return nil, NewDatabaseError("marshal to database structure failed", err)
	}

	_, err = r.cli.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("GameSystem"),
		Item:      item,
	})
	if err != nil {
		return nil, NewDatabaseError("failed to save the game", err)
	}

	// Create the domain output to be returned
	res := domain.NewGame(gameId, numRounds, createdAtAdm, closesAt, domain.OnGoing, admItem.Username)
	return &res, nil
}
