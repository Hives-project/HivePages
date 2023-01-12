package rest

import (
	"net/http"

	"github.com/Hives-project/HivePages/pkg/http/rest/handlers"
)

func (s *HttpServer) routes() {
	s.Router.PathPrefix("/pages/{uuid}").HandlerFunc(handlers.GetPageById(s.PageService)).Methods(http.MethodGet)
}
