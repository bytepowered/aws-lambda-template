package main

import (
    "aws-lambda-account/handler"
    aws "github.com/aws/aws-lambda-go/lambda"
    "github.com/bytepowered/aws-lambda-sdk/lambda"
)

func main() {
    lambda.SetupEnv()
    aws.StartWithOptions(handler.Authorize)
}
