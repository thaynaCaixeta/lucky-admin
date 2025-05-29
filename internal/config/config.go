package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	// Server config default values and keys
	defaultServerPort = "8080"
	serverPortKey     = "SERVER_PORT"

	defaultAddr = "localhost"
	addrKey     = "SERVER_ADDR"

	// DynamoDB config default values and keys
	localEndpointKey     = "DYNAMO_LOCAL_ENDPOINT"
	defaultLocalEndpoint = "http://localhost:8000"

	localRegionKey     = "DYNAMO_LOCAL_REGION"
	defaultLocalRegion = "us-west-2"

	tableNameKey     = "DYNAMO_TABLE_NAME"
	defaultTableName = "GameSystem"

	awsAccessKeyId_Key    = "AWS_ACCESS_KEY_ID"
	defaultAwsAccessKeyId = "your_key"

	awsSecretAccessKey_Key    = "AWS_SECRET_ACCESS_KEY"
	defaultAwsSecretAccessKey = "your_secret"

	awsRegionKey        = "AWS_REGION"
	defaultAwsRegionKey = "us-east-1"

	useLocalDB        = "USE_LOCAL_DYNAMO"
	defaultUseLocalDB = true
)

type AppConfig struct {
	DynamoConfig DynamoDBConfig
	ServerConfig ServerConfig
}

type DynamoDBConfig struct {
	LocalEndpoint      string
	LocalRegion        string
	TableName          string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	AwsRegion          string
	UseLocalDB         bool
}

type ServerConfig struct {
	Addr string
	Port string
}

func NewAppConfig() AppConfig {
	loadEnvConfig()

	return AppConfig{
		DynamoConfig: DynamoDBConfig{
			LocalEndpoint:      getEnvOrDefaultAsString(defaultLocalEndpoint, localEndpointKey),
			LocalRegion:        getEnvOrDefaultAsString(defaultLocalRegion, localRegionKey),
			TableName:          getEnvOrDefaultAsString(defaultTableName, tableNameKey),
			AwsAccessKeyId:     getEnvOrDefaultAsString(defaultAwsAccessKeyId, awsAccessKeyId_Key),
			AwsSecretAccessKey: getEnvOrDefaultAsString(defaultAwsSecretAccessKey, awsSecretAccessKey_Key),
			AwsRegion:          getEnvOrDefaultAsString(defaultAwsRegionKey, awsRegionKey),
			UseLocalDB:         getEnvOrDefaultAsBool(defaultUseLocalDB, useLocalDB),
		},
		ServerConfig: ServerConfig{
			Addr: getEnvOrDefaultAsString(defaultAddr, addrKey),
			Port: getEnvOrDefaultAsString(defaultServerPort, serverPortKey),
		},
	}
}

func loadEnvConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}
}

func getEnvOrDefaultAsString(defValue, key string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		return defValue
	}
	return envValue
}

func getEnvOrDefaultAsBool(defValue bool, key string) bool {
	envValue := os.Getenv(key)
	boolValue, err := strconv.ParseBool(key)
	if envValue == "" || err != nil {
		return defValue
	}
	return boolValue
}
