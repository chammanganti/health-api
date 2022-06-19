package awsp

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Loads the default AWS config
func LoadConfig() (aws.Config, error) {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return aws.Config{}, err
	}

	return config, nil
}
