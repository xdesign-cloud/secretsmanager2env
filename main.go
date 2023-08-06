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
		line := fmt.Sprintf(
			"export %s=%s",
			sanitizeKey(key),
			sanitizeValue(val),
		)
		line = sanitizeLine(line)

		output = append(output, line)
	}
	return output
}

// sanitizeLine will remove known dangerous characters from the entire line
func sanitizeLine(line string) string {
	return line
}

// santizeKey will format the key in a format valid for environment variables
func sanitizeKey(key string) string {
	key = strings.ToUpper(key)
	key = strings.ReplaceAll(key, " ", "_")
	key = strings.ReplaceAll(key, "-", "_")

	return key
}

// sanitizeValue will perform any sanitization which cannot happen at a line level
func sanitizeValue(value string) string {
	// Always quote the value - this handles spaces, but also a lot of bad characters
	value = fmt.Sprintf("\"%s\"", value)

	return value
}
