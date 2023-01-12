package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Hives-project/HivePages/pkg/config"
	"github.com/Hives-project/HivePages/pkg/page"
	"github.com/Hives-project/HivePages/pkg/protobuf/pb"
	pageRepository "github.com/Hives-project/HivePages/pkg/storage/mysql/page"
	"google.golang.org/grpc"
)

type PageServer struct {
	pb.UnimplementedPageServiceServer

	Server *grpc.Server

	environment string
	Cfg         config.GRPCConfig
	Sql         *sql.DB

	PageService page.PageService
}

const serverLog string = "[Server]: "

func NewPageServer(cfg *config.GRPCConfig, env string, version string, sql *sql.DB) *PageServer {
	baseUrl := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	newserver := &PageServer{
		environment: env,
		Cfg:         *cfg,
		Sql:         sql,
	}

	log.Println(serverLog+"started GRPc server on base url: ", baseUrl)

	return newserver
}

func (s *PageServer) Init() {
	projectRepository := pageRepository.NewPageRepository(s.Sql)
	s.PageService = page.NewPageService(projectRepository)
}

func (s *PageServer) Run(name string) {
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Cfg.Host, s.Cfg.Port))
		if err != nil {
			log.Fatalf("Failed to listen on port %s: %v", s.Cfg.Port, err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterPageServiceServer(grpcServer, s)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server over port %s: %v", s.Cfg.Port, err)
		}
	}()

	log.Println(serverLog+name, "is running..")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	log.Println(serverLog+name, "is shutting down..")

	os.Exit(0)
}

func (s *PageServer) GetPage(ctx context.Context, message *pb.PageRequest) (*pb.PageResponse, error) {
	page, err := s.PageService.GetPageById(context.Background(), message.Uuid)
	if err != nil {
		return nil, err
	}

	return &pb.PageResponse{Uuid: page.Uuid, PageName: page.PageName, Description: page.Description}, nil
}
