package emails

import (
	"app/internal/models"
	"html/template"
	"log"

	"github.com/wneessen/go-mail"
)

type EmailServiceController struct {
	client *mail.Client
}

func NewEmailServiceController(data models.EmailClient) *EmailServiceController {
	opts := []mail.Option{
		mail.WithPort(data.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(data.Username), // user email
		mail.WithPassword(data.Password), // user password or app password
	}

	if data.UseSSL {
		opts = append(opts, mail.WithSSL())
	}

	if data.Timeout != 0 {
		opts = append(opts, mail.WithTimeout(data.Timeout))
	}

	if !data.UseSSL {
		opts = append(opts, mail.WithTLSPortPolicy(mail.TLSPolicy(data.Port)))
	}

	mailClient, err := mail.NewClient(data.Host, opts...)

	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}

	return &EmailServiceController{
		client: mailClient,
	}
}

func (esc *EmailServiceController) SendEmail(request models.SendEmailRequest) error {
	m := mail.NewMsg()
	if err := m.From(request.From); err != nil {
		return err
	}
	if err := m.To(request.To); err != nil {
		return err
	}
	m.Subject(request.Subject)

	if request.DocType == mail.TypeTextHTML.String() {
		tlp, err := esc.templateEmail(request.Body)

		if err != nil {
			return err
		}
		m.SetBodyHTMLTemplate(tlp, request.Data)
	} else {
		m.SetBodyString(mail.TypeTextPlain, request.Body)
	}

	if err := esc.client.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (esc *EmailServiceController) templateEmail(body string) (*template.Template, error) {
	template, err := template.New("email").Parse(body)
	if err != nil {
		return nil, err
	}
	return template, nil
}

func (esc *EmailServiceController) Workers(id int, jobStream <-chan models.SendEmailRequest, jobsResponse chan<- error) error {
	for j := range jobStream {
		log.Printf("worker %d: started job: sending email to %s", id, j.To)
		jobsResponse <- esc.SendEmail(j)
	}
	return nil
}
