package app

import (
	"database/sql"
	"go-simple-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	// username : root
	// password : password
	// host 	: localhost
	// port 	: 3306
	// database : go_restful_api
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
