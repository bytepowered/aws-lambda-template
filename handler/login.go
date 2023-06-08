package handler

import (
    "aws-lambda-account/pkg"
    "context"
    "github.com/aws/aws-lambda-go/events"
    "github.com/bytepowered/aws-lambda-sdk/lambda"
    "log"
)

func Login(ctx context.Context, req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
    login := LoginUserVO{}
    if err := lambda.ParseBody(req.Body, req.IsBase64Encoded, &login); err != nil {
        log.Printf("ERROR: parse body failed: %s", err)
        return pkg.SendInvalidArgs()
    }
    if !lambda.CheckNotEmpty(login.Email, login.Password) {
        log.Printf("WARN: empty body: %s", login)
        return pkg.SendInvalidArgs()
    }
    return lambda.SendOK(map[string]any{
        "action": "LOGIN",
        "token":  "your-jwt-token",
        "user": map[string]any{
            "userid": 10000,
            "email":  login.Email,
        },
    })
}

type LoginUserVO struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
