package handlers

import (
	"fmt"
	"indexed-mail-search/server/pkg/handlers/contracts"
	"net/http"
	"sync"

	"github.com/go-chi/render"
)

type IndexerHandler struct {
	indexerEmailService contracts.IIndexerEmail
}

func NewIndexerHandler(ies contracts.IIndexerEmail) *IndexerHandler {
	return &IndexerHandler{
		indexerEmailService: ies,
	}
}

func (ih *IndexerHandler) IndexEmails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hola estoy en el handler")
	emailUsers, err := ih.indexerEmailService.GetMailUsers()
	if err != nil {
		//NewErrResponse(w, r, http.StatusInternalServerError, err)
		fmt.Println("Hola estoy en el primer erro de handler " + err.Error())
		return
	}

	var wg sync.WaitGroup
	for _, emailUser := range emailUsers {
		wg.Add(1)
		go ih.indexEmailByUser(emailUser, &wg)
	}
	wg.Wait()

	render.Status(r, http.StatusNoContent)
}

func (ih *IndexerHandler) indexEmailByUser(userEmail string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("estoy preparando la indexacion de: " + userEmail)
	emails, err := ih.indexerEmailService.ProcessMailsByUser(userEmail)
	if err != nil {
		return
	}

	err = ih.indexerEmailService.IndexEmails(emails)
	if err != nil {
		return
	}

}
