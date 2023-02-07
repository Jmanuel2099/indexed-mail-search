package handlers

import (
	customerror "indexed-mail-search/server/pkg/custom_error"
	"indexed-mail-search/server/pkg/domain"
	"indexed-mail-search/server/pkg/handlers/contracts"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// IndexedSearchHAandler is the handler for the IndexedSearch requests
type IndexedSearchHAandler struct {
	indexedSearchService contracts.IIndexedSearch
}

// NewIndexedSearchHAandler works as the conntrucutor of the IndexedSearchHAandler struc
func NewIndexedSearchHAandler(iss contracts.IIndexedSearch) *IndexedSearchHAandler {
	return &IndexedSearchHAandler{
		indexedSearchService: iss,
	}
}

// SearchTermInEmailsResponse is the response for the SearchTermInEmails method
type SearchTermInEmailsResponse struct {
	Emails []domain.Email `json:"emails"`
}

// SearchTermInEmails is the method that searches for a term in the emails.
func (ish *IndexedSearchHAandler) SearchTermInEmails(w http.ResponseWriter, r *http.Request) {
	term := chi.URLParam(r, "term")
	if len(term) > 150 || len(term) < 1 {
		errMessage := "term invalid. Length must be between 1 and 150"
		customerror.NewCustomError(http.StatusBadRequest, errMessage).ErrorResponseHandling(w, r)
		return
	}

	emails, err := ish.indexedSearchService.SearchInIndexedEmails(term)
	if err != nil {
		customerror.NewCustomError(http.StatusInternalServerError, err.Error()).ErrorResponseHandling(w, r)
		return
	}
	response := &SearchTermInEmailsResponse{
		Emails: emails,
	}
	render.JSON(w, r, response)
}
