package contracts

import "indexed-mail-search/server/pkg/domain"

// IIndexerEmail is the contract that must implement the service that performs the indexing of documents
type IIndexerEmail interface {
	GetMailUsers() ([]string, error)
	ProcessMailsByUser(user string) ([]domain.Email, error)
	IndexEmails(records []domain.Email) error
}
