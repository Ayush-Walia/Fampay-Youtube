// Package storage implements postgres connection and queries.
package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Ayush-Walia/Fampay-Youtube/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/slog"
)

func Init(conf *config.AppConfig) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.DBUser, conf.DBPass, conf.DBAddr, conf.DBName)
	var err error
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		slog.Fatal(err)
	}

	// Set max idle connection time to 5 minutes.
	db.SetConnMaxIdleTime(5 * time.Minute)
	if err = db.Ping(); err != nil {
		slog.Fatal(err)
	}
	slog.Info("Connected to DB successfully!")
}
