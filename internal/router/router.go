package router

import (
	handler "app/internal/email"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	handler *handler.Handler
}

func NewRouter(h *handler.Handler) *Router {
	return &Router{
		handler: h,
	}
}

func (r *Router) HandleRequest(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Default().Println("Received request:", request)

	switch request.Path {
	case "/emailSenderService/send-email":
		return r.handler.SendEmail(request)
	case "/emailSenderService/send-email/batch":
		return r.handler.SendEmailBatch(request)
	default:
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Email Service is running",
		}, nil
	}
}
