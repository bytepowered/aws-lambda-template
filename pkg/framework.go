package pkg

import (
    "github.com/aws/aws-lambda-go/events"
    "github.com/bytepowered/aws-lambda-sdk/lambda"
)

func SendInvalidArgs() (*events.APIGatewayV2HTTPResponse, error) {
    return lambda.SendERR("请求校验 : 参数不正确", 400)
}

func SendRemoteError() (*events.APIGatewayV2HTTPResponse, error) {
    return lambda.SendERR("服务繁忙 : 内部服务异常", 502)
}
