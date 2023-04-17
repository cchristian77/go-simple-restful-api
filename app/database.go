package app

import (
	"database/sql"
	"go-simple-restful-api/helper"
	"time"
)

/*
CREATE TABLE category(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL
) ENGINE = innoDB;
*/

func NewDB() *sql.DB {
	// username : root
	// password : password
	// host 	: localhost
	// port 	: 3306
	// database : go_restful_api
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
