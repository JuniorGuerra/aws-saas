package main

import (
	"app/cmd/emails_sender_service/config"

	handler "app/internal/email"
	"app/internal/models"
	"app/internal/router"
	service "app/services/emails"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

var r *router.Router

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

	h := handler.NewEmailHandler(emailServiceController)
	r = router.NewRouter(h)
}

func main() {
	fmt.Println("Starting Lambda function...")
	lambda.Start(r.HandleRequest)
}
