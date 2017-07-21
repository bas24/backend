package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func connectDB() error {
	client, err := sqlx.Connect("postgres", dbCockroachPathConst)
	if err != nil {
		return err
	}
	client.SetMaxIdleConns(12)
	client.SetMaxOpenConns(48)
	db = client

	return nil
}
