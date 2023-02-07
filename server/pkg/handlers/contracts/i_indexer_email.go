package contracts

import "indexed-mail-search/server/pkg/domain"

type IIndexerEmail interface {
	GetMailUsers() ([]string, error)
	ProcessMailsByUser(user string) ([]domain.Email, error)
	IndexEmails(records []domain.Email) error
}
