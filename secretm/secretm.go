package secretm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/inerhead/gossiouser/models"
)

// GetSecret is a function that returns a secret from AWS Secrets Manager
func GetSecret(cfg aws.Config, secretName string) (*models.SecretRDSJson, error) {
	client := secretsmanager.NewFromConfig(cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	}
	result, err := client.GetSecretValue(context.TODO(), input)
	if err != nil {
		fmt.Println("Error lectura secret OK !")
		return nil, err
	}
	var secret *models.SecretRDSJson
	err = json.Unmarshal([]byte(*result.SecretString), &secret)
	if err != nil {
		return nil, err
	}
	fmt.Println("lectura secret OK !")
	return secret, nil
}
