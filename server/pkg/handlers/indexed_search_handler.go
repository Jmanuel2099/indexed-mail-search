package handlers

import (
	"indexed-mail-search/server/pkg/handlers/contracts"
	"net/http"
)

type IndexedSearchHAandler struct {
	indexedSearchService contracts.IIndexedSearch
}

func NewIndexedSearchHAandler(iss contracts.IIndexedSearch) *IndexedSearchHAandler {
	return &IndexedSearchHAandler{
		indexedSearchService: iss,
	}
}

func (ish *IndexedSearchHAandler) SearchTermInEmails(w http.ResponseWriter, r *http.Request) {

}
