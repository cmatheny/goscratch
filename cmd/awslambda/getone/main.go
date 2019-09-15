package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cmatheny/goscratch/pkg/adapters/presenters/awslambda/getone"
)

func main() {
	handler := getone.NewHandler()
	lambda.Start(handler.Handle)
}
