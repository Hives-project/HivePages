package rest

import (
	"net/http"

	"github.com/Hives-project/HivePages/pkg/http/rest/handlers/page"
)

func (s *server) routes() {
	if s.environment != "develop" && s.environment != "production" {
		s.Router.Use(s.Cors)
	}

	// Example handlers subset
	exampleHandler := s.Router.PathPrefix("/pages").Subrouter()

	exampleHandler.HandleFunc("", page.CreatePageHandler(s.PageService)).Methods(http.MethodPost)
	exampleHandler.HandleFunc("", page.GetPageHandler(s.PageService)).Methods(http.MethodGet)

	exampleHandler.HandleFunc("/{uuid}", page.GetPageByUuidHandler(s.PageService)).Methods(http.MethodGet)
}
