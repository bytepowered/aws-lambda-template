package main

import (
	"aws-lambda-account/handler"
	aws "github.com/aws/aws-lambda-go/lambda"
	lambda "github.com/bytepowered/aws-lambda-sdk"
)

func main() {
	lambda.SetupEnv()
	aws.StartWithOptions(lambda.UseCORS(handler.Login))
}
