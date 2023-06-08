package main

import (
    "aws-lambda-account/handler"
    aws "github.com/aws/aws-lambda-go/lambda"
)

func main() {
    aws.StartWithOptions(handler.CORS)
}
