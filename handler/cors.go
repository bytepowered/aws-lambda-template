package handler

import (
    "context"
    "github.com/aws/aws-lambda-go/events"
)

func CORS(ctx context.Context, req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
    return &events.APIGatewayV2HTTPResponse{
        StatusCode: 200,
        Headers: map[string]string{
            "Access-Control-Allow-Origin":      "*",
            "Access-Control-Allow-Credentials": "true",
            "Access-Control-Allow-Methods":     "*",
            "Access-Control-Allow-Headers":     "Authorization",
        },
    }, nil
}
