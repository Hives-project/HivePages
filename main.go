package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Hives-project/HivePages/pkg/config"
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

	server := server.NewPageServer(&cfg.GRPC, cfg.Environment, cfg.Version, sql)
	server.Init()

	server.Run(cfg.Name)

	return nil
}
