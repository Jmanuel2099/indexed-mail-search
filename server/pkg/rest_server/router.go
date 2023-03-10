package restserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// configureRouter setups chi router mildelware for http communication
func (rs *RestServer) configureRouter() {
	rs.Router = chi.NewRouter()
	rs.Router.Use(middleware.RequestID)
	rs.Router.Use(middleware.Logger)
	rs.Router.Use(middleware.Recoverer)
	rs.Router.Use(middleware.URLFormat)
	rs.Router.Use(render.SetContentType(render.ContentTypeJSON))
	rs.configureCORS()
	rs.configureRoutes()
}

// configureCORS setups the router to allow cross-source resource sharing
// only with the VUE client and certain methods.
func (rs *RestServer) configureCORS() {
	rs.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5173"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

// configureRoutes sets up the routes for the app.
func (rs *RestServer) configureRoutes() {
	rs.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "Welcome to the Go Challenge!"})
	})

	rs.Router.Post("/indexer", rs.indexerHandler.IndexEmails)
	rs.Router.Get("/search", rs.indexedSearchHandler.SearchTermInEmails)
}
