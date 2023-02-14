package service

import (
	"encoding/json"
	"fmt"

	"indexed-mail-search/server/pkg/domain"
	contractsservice "indexed-mail-search/server/pkg/service/contracts_service"
)

const (
	defaultSearchType = "matchphrase"
	defaultMaxResults = 7
)

// IndexedSearchService is the struc that will communicate with the datasource
type IndexedSearchService struct {
	datasource contractsservice.IEmail
}

// NewIndexedSearchService works as the conntrucutor of the IndexedSearchService struc
func NewIndexedSearchService(ds contractsservice.IEmail) *IndexedSearchService {
	return &IndexedSearchService{
		datasource: ds,
	}
}

// SearchInIndexedEmails communicates with the zincsearch API to perform a search
// for a term in the indexed mails.
func (iss *IndexedSearchService) SearchInIndexedEmails(term string) ([]domain.Email, error) {
	// now := time.Now()
	// startTime := now.AddDate(0, 0, -7).Format("2006-01-02T15:04:05Z")
	// endTime := now.Format("2006-01-02T15:04:05Z")

	body := contractsservice.IndexedSearchRequest{
		SearchType: defaultSearchType,
		SortFields: []string{"-@timestamp"},
		From:       0,
		MaxResults: defaultMaxResults,
		Query: contractsservice.IndexedSearchRequestQuery{
			Term: term,
			// StartTime: startTime,
			// EndTime:   endTime,
		},
		// Source: []string{},
	}
	response, err := iss.datasource.IndexedSearch(body)
	if err != nil {
		return nil, err
	}
	fmt.Println(response.MaxScore)

	return mapResponseToEmails(response), nil
}

// mapResponseToEmails maps the IndexedSearchResponseResponse response structure to Emails
func mapResponseToEmails(response *contractsservice.IndexedSearchResponse) []domain.Email {
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
