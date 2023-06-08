package handler

import (
    "aws-lambda-account/pkg"
    "context"
    "github.com/aws/aws-lambda-go/events"
    "github.com/bytepowered/aws-lambda-sdk/lambda"
    "log"
)

const (
    _authorizeUserId   = "authorizeUserId"
    _authorizeUserRole = "authorizeUserRole"
)

func Authorize(ctx context.Context, req events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
    // https://docs.aws.amazon.com/zh_cn/apigateway/latest/developerguide/http-api-lambda-authorizer.html
    tokenStr := lambda.HeaderAuthorization(&events.APIGatewayV2HTTPRequest{
        Headers: req.Headers,
    })
    if tokenStr == "" || len(tokenStr) < 108 {
        log.Printf("WARN: empty/short token from header: %s", tokenStr)
        return events.APIGatewayV2CustomAuthorizerSimpleResponse{
            IsAuthorized: false,
        }, nil
    }
    log.Printf("INFO: verify token %s", tokenStr)
    token, err := pkg.AuthorizeParseToken(tokenStr)
    if err != nil {
        log.Printf("WARN: parse token error: %s, %s", tokenStr, err)
        return events.APIGatewayV2CustomAuthorizerSimpleResponse{
            IsAuthorized: false,
        }, nil
    }
    unwrap := func(v any, err error) any {
        return v
    }
    claims := token.Claims
    values := map[string]any{
        _authorizeUserId:   unwrap(claims.GetIssuer()).(string),
        _authorizeUserRole: unwrap(claims.GetSubject()).(string),
    }
    log.Println("INFO: verified token: isAuthorized=true")
    return events.APIGatewayV2CustomAuthorizerSimpleResponse{
        IsAuthorized: true, Context: values,
    }, nil
}
