package service_test

import (
	"app/config"
	"app/internal/models"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmailServiceController_SendEmail(t *testing.T) {
	conf := config.LoadConfig()
	t.Run("Text_Plain_Test", func(t *testing.T) {
		emailServiceController := service.NewEmailServiceController(models.EmailClient{
			Host:     conf.EmailClient.Host,
			Port:     conf.EmailClient.Port,
			Username: conf.EmailClient.Username,
			Password: conf.EmailClient.Password,
			Timeout:  conf.EmailClient.Timeout,
			UseSSL:   conf.EmailClient.UseSSL,
		})

		err := emailServiceController.SendEmail(models.SendEmailRequest{
			SendEmailRequestCommon: models.SendEmailRequestCommon{
				Subject: "Test Email from Go",
				Body:    "<h1>This is a test email sent from Go using go-mail library.</h1>",
				From:    conf.EmailClient.Username,
			},
			SendEmailRequestIndividualCommon: models.SendEmailRequestIndividualCommon{
				To: "viwimeg969@inupup.com",
			},

			// DocType: "text/html",
			// Data: nil,
		})

		require.NoError(t, err)
	})

	t.Run("Text_HTML_Test", func(t *testing.T) {
		emailServiceController := service.NewEmailServiceController(models.EmailClient{
			Host:     conf.EmailClient.Host,
			Port:     conf.EmailClient.Port,
			Username: conf.EmailClient.Username,
			Password: conf.EmailClient.Password,
			Timeout:  conf.EmailClient.Timeout,
			UseSSL:   conf.EmailClient.UseSSL,
		})

		err := emailServiceController.SendEmail(models.SendEmailRequest{
			SendEmailRequestCommon: models.SendEmailRequestCommon{
				Subject: "Test Email from Go",
				From:    conf.EmailClient.Username,
				Body:    "<h1>This is a test email sent from Go using go-mail library {{.name}}.</h1>",
				DocType: "text/html",
			},
			SendEmailRequestIndividualCommon: models.SendEmailRequestIndividualCommon{
				To: "viwimeg969@inupup.com",
				Data: map[string]string{
					"name": "John Doe",
				},
			},
		})

		require.NoError(t, err)
	})
}
