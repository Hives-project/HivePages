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

type server struct {
	environment string
	Cfg         config.HTTPConfig

	Server *http.Server
	Router *mux.Router
	Sql    *sql.DB

	PageService page.PageService
}

const serverLog string = "[Server]: "

func NewServer(cfg *config.HTTPConfig, env string, version string, sql *sql.DB) *server {
	baseUrl := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	s := &server{
		environment: env,
		Cfg:         *cfg,
		Server: &http.Server{
			Addr:         baseUrl,
			WriteTimeout: cfg.WriteTimeOut,
			ReadTimeout:  cfg.ReadTimeOut,
			IdleTimeout:  cfg.IdleTimeOut,
		},
		Sql:    sql,
		Router: mux.NewRouter(),
	}

	log.Println(serverLog+"started api on base url: ", baseUrl)

	// Generic routes
	s.Router.NotFoundHandler = http.HandlerFunc(handleNotFound)

	s.Server.Handler = s.Router

	return s
}

func (s *server) Init() {
	projectRepository := pageRepository.NewPageRepository(s.Sql)
	s.PageService = page.NewPageService(projectRepository)

	s.routes()
}

func (s *server) Run(name string) {
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

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("404 - Endpoint was not found")
	handlers.RenderErrorResponse(w, err.Error(), r.URL.Path, err)
}
