package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, ev struct{}) (string, error) {
	return "", errors.New("FAIL")
}

func main() {
	lambda.Start(HandleRequest)
}
