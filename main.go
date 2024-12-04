package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/inerhead/gossiouser/awsgo"
	"github.com/inerhead/gossiouser/bd"
	"github.com/inerhead/gossiouser/models"
)

func main() {
	lambda.Start(handlerLambda)
}

func handlerLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (*events.CognitoEventUserPoolsPostConfirmation, error) {
	// Your code here
	conf := awsgo.ConfigAWS()
	if !ValidarParametros() {
		fmt.Println("No se encontraron los parametros necesarios, enviar secretManager")
		err := errors.New("No se encontraron los parametros necesarios, enviar secretManager")
		return nil, err
	}
	var datos models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email: ", att)
		case "sub":
			datos.UserUUID = att
			fmt.Println("UUID: ", att)
		}

	}

	_, errs := bd.ReadSecret(conf)
	if errs != nil {
		fmt.Println("Error lectura secret !")
		return nil, errs
	}
	err := bd.SignUp(datos)
	if err != nil {
		fmt.Println("Error en el registro", err.Error())
		return nil, err
	}
	return &event, nil
}

func ValidarParametros() bool {
	_, traerParametro := os.LookupEnv("SecretName")
	return traerParametro
}
