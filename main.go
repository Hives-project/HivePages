package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Hives-project/HivePages/pkg/config"
	"github.com/Hives-project/HivePages/pkg/http/rest"
	"github.com/Hives-project/HivePages/pkg/kafka/consumer"
	"github.com/Hives-project/HivePages/pkg/kafka/producer"
	"github.com/Hives-project/HivePages/pkg/protobuf/server"
	"github.com/Hives-project/HivePages/pkg/storage/mysql"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.NewConfig()
	err := cfg.LoadConfig()
	if err != nil {
		return errors.New(err.Error())
	}

	sql, err := mysql.Connect(cfg.Sql)
	if err != nil {
		return err
	}
	httpServer := rest.NewHTTPServer(cfg, cfg.Environment, cfg.Version, sql)
	httpServer.Init()

	server := server.NewPageServer(&cfg.GRPC, cfg.Environment, cfg.Version, sql)
	server.Init()

	producer.Init(cfg.Kafka)

	go consumer.StartKafkaConsumer(cfg.Kafka, server.PageService)

	go httpServer.Run(cfg.Name)

	server.Run(cfg.Name)

	return nil
}
