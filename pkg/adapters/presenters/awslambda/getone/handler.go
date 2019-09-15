package getone

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/cmatheny/goscratch/pkg/domain/user"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

type Handler struct{}

func (h *Handler) Handle(request Request) (Response, error) {
	log.Println(request)
	if request.Headers["Fail"] == "true" {
		return Response{}, fmt.Errorf("fail header")
	}

	u := user.User{
		ID:   "123",
		Name: "Bob",
	}

	data, err := json.Marshal(u)
	var response Response
	if err != nil {
		response = Response{
			StatusCode: 500,
		}
	}

	response = Response{
		StatusCode:      200,
		Body:            string(data),
		IsBase64Encoded: false,
	}

	return response, nil
}

func NewHandler() *Handler {
	return &Handler{}
}
