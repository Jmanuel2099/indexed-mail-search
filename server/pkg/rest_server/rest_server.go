package restserver

import (
	"fmt"
	"net/http"

	"indexed-mail-search/server/pkg/datasource"
	"indexed-mail-search/server/pkg/handlers"
	"indexed-mail-search/server/pkg/service"

	"github.com/go-chi/chi/v5"
)

const (
	serverPort = ":8000"
)

// RestServer contains the structure of the server
type RestServer struct {
	Router               *chi.Mux
	httpClient           *http.Client
	indexerHandler       *handlers.IndexerHandler
	indexedSearchHandler *handlers.IndexedSearchHAandler
}

// NewRestServer works as the conntrucutor of the RestServer struc
func NewRestServer() *RestServer {
	server := &RestServer{}
	server.configureHttpClient()
	server.configureDependencies()
	server.configureRouter()

	return server
}

// configureHttpClient setups the http client which needs the sling client
func (rs *RestServer) configureHttpClient() {
	rs.httpClient = &http.Client{}
}

// configureHandlers setups dependency injection
func (rs *RestServer) configureDependencies() {
	datasourceZincSearch := datasource.NewZincsearchClient(rs.httpClient)

	indexEmailService := service.NewIndexerService(datasourceZincSearch)
	indexerHandler := handlers.NewIndexerHandler(indexEmailService)
	rs.indexerHandler = indexerHandler

	indexedSearchService := service.NewIndexedSearchService(datasourceZincSearch)
	indexedSearchHandler := handlers.NewIndexedSearchHAandler(indexedSearchService)
	rs.indexedSearchHandler = indexedSearchHandler
}

func (rs *RestServer) RunServer() {
	fmt.Println("Application is running on host http://localhost:8000/")

	err := http.ListenAndServe(serverPort, rs.Router)
	if err != nil {
		panic(err)
	}
}
