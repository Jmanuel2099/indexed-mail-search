package restserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func (rs *RestServer) configureRouter() {
	rs.Router = chi.NewRouter()
	rs.Router.Use(middleware.RequestID)
	rs.Router.Use(middleware.Logger)
	rs.Router.Use(middleware.Recoverer)
	rs.Router.Use(middleware.URLFormat)
	rs.Router.Use(render.SetContentType(render.ContentTypeJSON))
	rs.configureRoutes()
}

func (rs *RestServer) configureRoutes() {
	rs.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "Welcome to the Go Challenge!"})
	})

	rs.Router.Post("/indexer", rs.indexerHandler.IndexEmails)
	rs.Router.Get("/search", rs.indexedSearchHandler.SearchTermInEmails)
}
