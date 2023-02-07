package contracts

type IEmail interface {
	CreateEmails(emails interface{}) (*CreateEmailsResponse, error)
	IndexedSearch(bodyrequest IndexedSearchRequest) (*IndexedSearchResponse, error)
	// DeleteEmails() error
}

type CreateEmailsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

type IndexedSearchRequest struct {
	SearchType string                    `json:"search_type"`
	SortFields []string                  `json:"sort_fields"`
	From       int                       `json:"from"`
	MaxResults int                       `json:"max_results"`
	Query      IndexedSearchRequestQuery `json:"query"`
	Source     map[string]interface{}    `json:"_source"`
}

type IndexedSearchRequestQuery struct {
	Term      string `json:"term"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type CreateEmailsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

type IndexedSearchResponse struct {
	Hits struct {
		Hits []struct {
			ID        string                 `json:"_id"`
			Timestamp string                 `json:"@timestamp"`
			Score     float64                `json:"_score"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}

type ErrorReponse struct {
	ErrorMessage string `json:"error"`
}
