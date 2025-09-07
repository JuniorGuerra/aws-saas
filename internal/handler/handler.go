package handler

import (
	"app/internal/models"
	"app/services/emails"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	EmailSvc *emails.EmailServiceController
}

func NewHandlerController(emailSvc *emails.EmailServiceController) *Handler {
	return &Handler{emailSvc}
}

func (h Handler) SendEmail(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var emailRequest models.SendEmailRequest
	err := json.Unmarshal([]byte(request.Body), &emailRequest)
	if err != nil {
		log.Println("Error unmarshalling request body:", err)
		return nil, err
	}

	err = h.EmailSvc.SendEmail(emailRequest)

	if err != nil {
		log.Println("Error sending email:", err)
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Email sent successfully",
	}, nil
}

func (h Handler) SendEmailBatch(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var emailBatchRequests models.SendEmailBatchRequest

	err := json.Unmarshal([]byte(request.Body), &emailBatchRequests)
	if err != nil {
		log.Println("Error unmarshalling request body:", err)
		return nil, err
	}

	jobsStream := make(chan models.SendEmailRequest, len(emailBatchRequests.ToList))
	jobsResponse := make(chan error, len(emailBatchRequests.ToList))

	// Start worker pool
	for id := range 5 { // change for less workers
		go h.EmailSvc.Workers(id, jobsStream, jobsResponse)
	}

	// Dispatch jobs
	for _, recipient := range emailBatchRequests.ToList {
		jobsStream <- models.SendEmailRequest{
			SendEmailRequestCommon: emailBatchRequests.SendEmailRequestCommon,
			SendEmailRequestIndividualCommon: models.SendEmailRequestIndividualCommon{
				To:   recipient.To,
				Data: recipient.Data,
			},
		}
	}

	close(jobsStream)

	for i := 0; i < len(emailBatchRequests.ToList); i++ {
		err := <-jobsResponse
		if err != nil {
			log.Println("Error sending email in batch:", err)
		}
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Batch emails processed",
	}, nil

}
