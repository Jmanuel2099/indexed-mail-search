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
	Source     []string                  `json:"_source"`
	// Source     map[string]interface{}    `json:"_source"`
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
	Took     float64 `json:"took"`
	TimedOut bool    `json:"timed_out"`
	MaxScore float64 `json:"max_score"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Index     string                 `json:"_index"`
			Type      string                 `json:"_type"`
			ID        string                 `json:"_id"`
			Score     float64                `json:"_score"`
			Timestamp string                 `json:"@timestamp"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type ErrorReponse struct {
	ErrorMessage string `json:"error"`
}
