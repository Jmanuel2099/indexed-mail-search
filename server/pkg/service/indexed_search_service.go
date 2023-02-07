package service

import (
	"encoding/json"
	"indexed-mail-search/server/pkg/domain"
	"indexed-mail-search/server/pkg/handlers/contracts"
	"time"
)

const (
	defaultSearchType = "matchphrase"
	defaultMaxResults = 7
)

// IndexedSearchService is the struc that will communicate with the datasource
type IndexedSearchService struct {
	datasource contracts.IEmail
}

// NewIndexedSearchService works as the conntrucutor of the IndexedSearchService struc
func NewIndexedSearchService(ds contracts.IEmail) *IndexedSearchService {
	return &IndexedSearchService{
		datasource: ds,
	}
}

// SearchInIndexedEmails communicates with the zincsearch API to perform a search
// for a term in the indexed mails.
func (iss *IndexedSearchService) SearchInIndexedEmails(term string) ([]domain.Email, error) {
	now := time.Now()
	startTime := now.AddDate(0, 0, -7).Format("2006-01-02T15:04:05Z")
	endTime := now.Format("2006-01-02T15:04:05Z")

	body := contracts.IndexedSearchRequest{
		SearchType: defaultSearchType,
		SortFields: []string{"-@timestamp"},
		From:       0,
		MaxResults: defaultMaxResults,
		Query: contracts.IndexedSearchRequestQuery{
			Term:      term,
			StartTime: startTime,
			EndTime:   endTime,
		},
		Source: []string{},
	}
	response, err := iss.datasource.IndexedSearch(body)
	if err != nil {

		return nil, err
	}

	return mapResponseToEmails(response), nil
}

// mapResponseToEmails maps the IndexedSearchResponseResponse response structure to Emails
func mapResponseToEmails(response *contracts.IndexedSearchResponse) []domain.Email {
	var emails []domain.Email

	for _, hit := range response.Hits.Hits {
		var email domain.Email
		contetEmialBytes, _ := json.Marshal(hit.Source)

		err := json.Unmarshal(contetEmialBytes, &email)
		if err != nil {
			continue
		}
		emails = append(emails, email)
	}

	return emails
}
