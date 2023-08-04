package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func main() {
	secretName := os.Args[1]
	appConfig := getFromSecretsManager(secretName)
	envVars := convertToEnvVarStatements(appConfig)

	for line := range envVars {
		fmt.Println(line)
	}
}

func getFromSecretsManager(secretName string) map[string]string {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	svc := secretsmanager.NewFromConfig(config)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}

	appConfig := map[string]string{}
	json.Unmarshal([]byte(*result.SecretString), &appConfig)

	return appConfig
}

func convertToEnvVarStatements(config map[string]string) []string {
	var output []string
	for key, val := range config {
		output = append(output, fmt.Sprintf("export %s=%s", strings.ToUpper(key), val))
	}
	return output
}
