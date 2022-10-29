package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Hives-project/HivePages/pkg/config"
)

// All Connection logic
// Think of functionallity like initializing the database connection on app start
// As well as getting a collection for MongoDB

func Connect(cfg config.MySQLConfig) (*sql.DB, error) {
	ctx := context.Background()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return nil, fmt.Errorf("failed database connection because: %s", err.Error())
	}

	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed ping because: %s", err.Error())
	}

	log.Println("MySQL database connected!")

	var version string

	err = db.QueryRowContext(ctx, "SELECT VERSION()").Scan(&version)
	if err != nil {
		return nil, fmt.Errorf("failed version query because: %s", err.Error())
	}

	log.Printf("MySQL database version %s\n", version)
	return db, nil
}
