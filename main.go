package main

import (
	"context"
	"fmt"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/punigonzalez/gambitUser/awsgo"
	"github.com/punigonzalez/gambitUser/bd"
	"github.com/punigonzalez/gambitUser/models"
)

func main() {
	lambda.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	if !ValidoParamtros() {
		fmt.Println("Error en los parametros, debe enviar 'SecretName'")
		os.Exit(0)
	}

	var datos models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Email = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el secret " + err.Error())
		return event, err
	}

	err = bd.DbConnect()
	if err != nil {
		fmt.Println("Error al conectar a la BD " + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)
	return event, err

}

func ValidoParamtros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
