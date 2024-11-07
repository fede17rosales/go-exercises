package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"twittergo/awsgo"
	"twittergo/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dateSecret models.Secret
	fmt.Println(" > Pido Secreto " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dateSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dateSecret)
	fmt.Println(" > Lectura de Secret OK" + secretName)
	return dateSecret, nil
}
