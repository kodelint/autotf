package commands

import (
	"fmt"
	"os"
)

func FetchEnvVariables() error {
	if os.Getenv(EnvS3Bucket) != "" {
		S3Bucket = os.Getenv("S3_BUCKET")
	} else {
		return fmt.Errorf("S3_BUCKET environment variable is not set")
	}
	if os.Getenv(EnvRegion) != "" {
		Region = os.Getenv("REGION")
	} else {
		return fmt.Errorf("REGION environment variable is not set")
	}
	if os.Getenv(EnvDynamoDB) != "" {
		DynamoDB = os.Getenv("DYNAMODB")
	} else {
		return fmt.Errorf("REGION environment variable is not set")
	}
	return nil
}

func GenerateBackendConfig() (BackendConfig, error) {
	var backendConfig BackendConfig
	if S3Bucket != "" {
		backendConfig.S3Bucket = S3Bucket
	} else {
		return BackendConfig{},fmt.Errorf("couldn't find the S3 Bucket, check S3_BUCKET environement variable")
	}
	if DynamoDB != "" {
		backendConfig.DynamoDB = DynamoDB
	} else {
		return BackendConfig{},fmt.Errorf("couldn't find the DynamoDB, check DYNAMODB environement variable")
	}
	if Region != "" {
		backendConfig.Region = Region
	} else {
		return BackendConfig{},fmt.Errorf("couldn't find the Region, check REGION environement variable")
	}
	if Key != "" {
		backendConfig.Key = Key
	} else {
		return BackendConfig{}, fmt.Errorf("couldn't find the Generated Key")
	}
	return backendConfig, nil
}