package datasource

type CreateEmailsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}
