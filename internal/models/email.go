package models

type SendEmailRequestCommon struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	DocType string `json:"doc_type"` // "text/plain" or "text/html"
}

type SendEmailRequestIndividualCommon struct {
	To   string `json:"to"`
	Data any    `json:"template_data"` // Data for templating
}

type SendEmailRequest struct {
	SendEmailRequestCommon
	SendEmailRequestIndividualCommon
}

type SendEmailBatchRequest struct {
	SendEmailRequestCommon
	ToList []SendEmailRequestIndividualCommon `json:"to_list"`
}
