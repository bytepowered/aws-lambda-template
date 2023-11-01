package pkg

import (
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/bytepowered/aws-lambda-sdk"
)

func SendInvalidArgs() (*events.APIGatewayV2HTTPResponse, error) {
	return lambda.SendERR("请求校验 : 参数不正确", 400)
}
