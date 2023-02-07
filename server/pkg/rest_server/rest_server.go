package restserver

import (
	"indexed-mail-search/server/pkg/datasource"
	"indexed-mail-search/server/pkg/handlers"
	"indexed-mail-search/server/pkg/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	serverPort = ":8000"
)

type RestServer struct {
	Router               *chi.Mux
	httpClient           *http.Client
	indexerHandler       *handlers.IndexerHandler
	indexedSearchHandler *handlers.IndexedSearchHAandler
}

func NewRestServer() *RestServer {
	server := &RestServer{}
	server.configureHttpClient()
	server.configureHandlers()
	server.configureRouter()

	return server
}

func (rs *RestServer) configureHttpClient() {
	rs.httpClient = &http.Client{}
}

func (rs *RestServer) configureHandlers() {
	datasourceZincSearch := datasource.NewZincsearchClient(rs.httpClient)

	indexEmailService := service.NewIndexerService(datasourceZincSearch)
	indexerHandler := handlers.NewIndexerHandler(indexEmailService)
	rs.indexerHandler = indexerHandler

	indexedSearchService := service.NewIndexedSearchService(datasourceZincSearch)
	indexedSearchHandler := handlers.NewIndexedSearchHAandler(indexedSearchService)
	rs.indexedSearchHandler = indexedSearchHandler
}

func (rs *RestServer) RunServer() {
	err := http.ListenAndServe(serverPort, rs.Router)
	if err != nil {
		panic(err)
	}
}
