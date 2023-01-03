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
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", conf.DBUser, conf.DBPass, conf.DBAddr, conf.DBName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		slog.Fatal(err)
	}

	db.SetConnMaxIdleTime(time.Minute * 5)
	if err = db.Ping(); err != nil {
		slog.Fatal(err)
	}
	slog.Info("Connected to DB successfully!")

	VideosDao = newVideosDAO(db)
}
