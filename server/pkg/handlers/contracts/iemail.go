package contracts

type IEmail interface {
	CreateEmails(emails interface{}) (*CreateEmailsResponse, error)
	// DeleteEmails() error
	// SearchEmails() error
}

type CreateEmailsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

type ErrorReponse struct {
	ErrorMessage string `json:"error"`
}
