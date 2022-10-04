package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Hives-project/HivePages/pkg/config"
	"github.com/Hives-project/HivePages/pkg/http/rest"
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

	server := rest.NewServer(&cfg.HTTP, cfg.Environment, cfg.Version, sql)
	server.Init()

	// Runs the new server instance.
	server.Run(cfg.Name)

	return nil
}
