package secretM

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/punigonzalez/gambitUser/awsgo"
	"github.com/punigonzalez/gambitUser/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSjson, error) {
	var datosSecret models.SecretRDSjson
	fmt.Println(" > Pido Secreto " + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura Secret OK " + nombreSecret)
	return datosSecret, nil
}
