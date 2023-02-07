package contracts

import "indexed-mail-search/server/pkg/domain"

type IIndexedSearch interface {
	SearchInIndexedEmails(term string) ([]domain.Email, error)
}
