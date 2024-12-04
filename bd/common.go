package bd

import (
	"database/sql"

	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/inerhead/gossiouser/models"
	"github.com/inerhead/gossiouser/secretm"
)

var secret *models.SecretRDSJson
var DB *sql.DB

func ReadSecret(cfg aws.Config) (*models.SecretRDSJson, error) {
	useSecretManager, err := strconv.ParseBool(os.Getenv("USE_SECRET_MANAGER"))
	if err != nil {
		fmt.Println("Error en la variable de entorno USE_SECRET_MANAGER")
		return nil, err
	}
	secretName := os.Getenv("SecretName")

	if secret == nil {
		var err error
		if useSecretManager {
			secret, err = secretm.GetSecret(cfg, secretName)

			if err != nil {
				fmt.Println("Error lectura secret !")
				return nil, err
			}
		} else {
			secret = &models.SecretRDSJson{
				Username:            "root",
				Password:            "Agosto2013*.",
				Engine:              "MySQL Community",
				Host:                "gobd.c3amu4iicnlu.us-east-1.rds.amazonaws.com",
				Port:                3306,
				DbClusterIdentifier: "gobd",
				DBName:              "gossio",
			}
		}
	}
	return secret, nil
}

func DbConnect() error {
	var err error
	if DB == nil {

		DB, err = sql.Open(secret.Engine, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			secret.Host, secret.Port, secret.Username, secret.Password, secret.DBName))
		if err != nil {
			fmt.Println("Error en la conexion a la base de datos", err.Error())
			return err
		}
		err = DB.Ping()
		if err != nil {
			fmt.Println("Error en el ping a la base de datos", err.Error())
			return err
		}
	}
	fmt.Println("Conexion a la base de datos exitosa")
	return nil
}
