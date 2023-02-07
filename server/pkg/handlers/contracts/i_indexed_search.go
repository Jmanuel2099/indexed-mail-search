package contracts

import "indexed-mail-search/server/pkg/domain"

type IIndexedSearch interface {
	SearchInIndexedEmails(indexName string, term string) ([]domain.Email, error)
}
