package bd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/inerhead/gossiouser/models"
	"github.com/inerhead/gossiouser/secretm"
)

var secret *models.SecretRDSJson

func ReadSecret(cfg aws.Config) (*models.SecretRDSJson, error) {
	secretName := os.Getenv("SecretName")

	if secret == nil {
		var err error
		secret, err = secretm.GetSecret(cfg, secretName)

		if err != nil {
			fmt.Println("Error lectura secret !")
			return nil, err
		}
	}
	return secret, nil
}
