package rest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Hives-project/HivePages/pkg/config"
	"github.com/Hives-project/HivePages/pkg/http/rest/handlers"
	"github.com/Hives-project/HivePages/pkg/page"
	pageRepository "github.com/Hives-project/HivePages/pkg/storage/mysql/page"
	"github.com/gorilla/mux"
)

type HttpServer struct {
	environment string
	HTTPcfg     config.HTTPConfig
	Server      *http.Server
	Router      *mux.Router
	Sql         *sql.DB

	PageService page.PageService
}

const serverLog string = "[HTTP Server]: "

func NewHTTPServer(cfg *config.Config, env string, version string, sql *sql.DB) *HttpServer {
	s := &HttpServer{
		environment: env,
		Server: &http.Server{
			Addr:         cfg.HTTP.PageUrl,
			WriteTimeout: cfg.HTTP.WriteTimeOut,
			ReadTimeout:  cfg.HTTP.ReadTimeOut,
			IdleTimeout:  cfg.HTTP.IdleTimeOut,
		},
		Sql:    sql,
		Router: mux.NewRouter(),
	}

	log.Println(serverLog+"started http server on base url: ", cfg.HTTP.PageUrl)

	// Generic routes
	s.Router.NotFoundHandler = http.HandlerFunc(handleNotFound)

	s.Server.Handler = s.Router

	return s
}

func (s *HttpServer) Init() {
	projectRepository := pageRepository.NewPageRepository(s.Sql)
	s.PageService = page.NewPageService(projectRepository)

	s.routes()
}

func (s *HttpServer) Run(name string) {
	var wait time.Duration

	s.Server.Handler = s.Router

	go func() {
		if err := s.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Println(serverLog+name, "is running..")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	s.Server.Shutdown(ctx)

	log.Println(serverLog+name, "is shutting down..")

	os.Exit(0)
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("404 - Endpoint was not found")
	handlers.RenderErrorResponse(w, err.Error(), r.URL.Path, err)
}
