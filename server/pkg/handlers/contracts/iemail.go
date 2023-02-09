package contracts

// IEmail is the contract that must implement the data source.
type IEmail interface {
	CreateEmails(emails interface{}) (*CreateEmailsResponse, error)
	IndexedSearch(bodyrequest IndexedSearchRequest) (*IndexedSearchResponse, error)
}

// CreateEmailsRequest is the request for the CreateEmails function
type CreateEmailsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

// IndexedSearchRequest is the request for the IndexedSearch function
type IndexedSearchRequest struct {
	SearchType string                    `json:"search_type"`
	SortFields []string                  `json:"sort_fields"`
	From       int                       `json:"from"`
	MaxResults int                       `json:"max_results"`
	Query      IndexedSearchRequestQuery `json:"query"`
	// Source     []string                  `json:"_source"`
}

// IndexedSearchRequestQuery is the query for the IndexedSearch function
type IndexedSearchRequestQuery struct {
	Term      string `json:"term"`
	// StartTime string `json:"start_time"`
	// EndTime   string `json:"end_time"`
}

// CreateEmailsResponse is the response for the CreateEmailsResponse function
type CreateEmailsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

// IndexedSearchResponse is the response for the IndexedSearch function
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

// ErrorReponse is the response of datasource when an error occurs
type ErrorReponse struct {
	ErrorMessage string `json:"error"`
}
