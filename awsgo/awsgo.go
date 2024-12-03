package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// ConfigAWS is a function that returns a configuration for the AWS SDK
func ConfigAWS() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	return cfg
}
