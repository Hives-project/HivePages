package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

	// if err := run(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// 	os.Exit(1)
	// }
}

// func run() error {
// 	cfg := config.NewConfig()

// 	err := cfg.LoadConfig()
// 	if err != nil {
// 		return errors.New(err.Error())
// 	}

// 	sql, err := mysql.Connect(cfg.Sql)
// 	if err != nil {
// 		return err
// 	}

// 	server := rest.NewServer(&cfg.HTTP, cfg.Environment, cfg.Version, sql)
// 	server.Init()

// 	// Runs the new server instance.
// 	server.Run(cfg.Name)

// 	return nil
// }
