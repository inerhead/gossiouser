package bd

import (
	"fmt"

	"github.com/inerhead/gossiouser/models"
	"github.com/inerhead/gossiouser/tools"
)

func SignUp(sign models.SignUp) error {
	fmt.Println("Comienza registro")

	err := DbConnect()
	if err != nil {
		fmt.Println("Error en la conexion a la base de datos 2", err.Error())
		return err
	}
	defer DB.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sign.UserEmail + "', '" + sign.UserUUID + "', , '" + tools.GetTime() + "')"
	fmt.Println("Formato sentencia", sentencia)
	_, err = DB.Exec(sentencia)
	if err != nil {
		fmt.Println("Error en la insercion de datos", err.Error())
	}
	fmt.Println("Registro exitoso")
	return nil

}
