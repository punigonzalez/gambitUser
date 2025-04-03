package main

import (
	"context"
	"fmt"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/punigonzalez/gambitUser/awsgo"
)

func main() {
	lambda.Start(ejecutoLambda)

}

func ejecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	if !ValidoParamtros() {
		fmt.Println("Error en los parametros, debe enviar 'SecretName'")
		os.Exit(0)
	}

}

func ValidoParamtros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
