package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Jack5758/internal/config"
	"github.com/Jack5758/pkg/util/timeparser"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDB(conf config.DBConfig) (*sql.DB, error) {
	var (
		driver   = conf.Driver
		user     = conf.User
		password = conf.Password
		database = conf.DB
		timeout  = conf.Timeout
	)
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, database)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("data source name is not correct %w", err)
	}
	err = testConnection(db, timeout)
	if err != nil {
		return nil, err

	}
	return db, nil
}

func testConnection(db *sql.DB, timeout string) error {
	ctx := context.Background()
	duration, err := timeparser.ParseDuration(timeout)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("connection timeout at startup %w", err)
	}
	return nil
}
