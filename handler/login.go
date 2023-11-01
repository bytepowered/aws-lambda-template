package handler

import (
	"aws-lambda-account/pkg"
	"context"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/bytepowered/aws-lambda-sdk"
	"log"
)

func Login(ctx context.Context, req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	args := LoginUserVO{}
	if err := lambda.ParseBody(req.Body, req.IsBase64Encoded, &args); err != nil {
		log.Printf("ERROR: parse body failed: %s", err)
		return pkg.SendInvalidArgs()
	}
	if !lambda.CheckNotEmpty(args.Email, args.Password) {
		log.Printf("WARN: empty body: %s", args)
		return pkg.SendInvalidArgs()
	}
	token, err := lambda.JWTSigned(args.Email, "000001")
	if err != nil {
		log.Printf("ERROR: sign jwt error: %s", err)
		return pkg.SendInvalidArgs()
	}
	return lambda.SendOK(map[string]any{
		"action": "LOGIN",
		"token":  token,
		"user": map[string]any{
			"userid": "000001",
			"email":  args.Email,
		},
	})
}

type LoginUserVO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
