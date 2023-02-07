package contracts

import "indexed-mail-search/server/pkg/domain"

// IIndexedSearch is the contract to be implemented by the service that performs the indexed search
type IIndexedSearch interface {
	SearchInIndexedEmails(term string) ([]domain.Email, error)
}
