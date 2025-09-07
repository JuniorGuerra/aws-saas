package main

import (
	"app/config"
	"app/internal/handler"
	"app/internal/models"
	"app/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var h *handler.Handler

func init() {
	conf := config.LoadConfig()

	emailServiceController := service.NewEmailServiceController(models.EmailClient{
		Host:     conf.EmailClient.Host,
		Port:     conf.EmailClient.Port,
		Username: conf.EmailClient.Username,
		Password: conf.EmailClient.Password,
		Timeout:  conf.EmailClient.Timeout,
		UseSSL:   conf.EmailClient.UseSSL,
	})

	h = handler.NewHandlerController(emailServiceController)
}

func main() {
	fmt.Println("Starting Lambda function...")
	lambda.Start(lambdaHandler)
}

func lambdaHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	log.Default().Println("Received request:", request)

	switch request.Path {
	case "/emailSenderService/send-email":
		return h.SendEmail(request)
	case "/emailSenderService/send-email/batch":
		return h.SendEmailBatch(request)
	default:
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Email Service is running",
		}, nil
	}
}
