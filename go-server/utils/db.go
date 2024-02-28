package utils

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func init() {
	db, err := sqlx.Connect("sqlite3", DbPath)
	if err != nil {
		log.Fatal("Connect DB error: " + err.Error())
	}
	log.Println("Check tables...")

	sqlUser := "create table if not exists user(uid text, account text, password text, permission int, register_time text, token text)"
	_, createUserTableErr := db.Exec(sqlUser)

	if createUserTableErr != nil {
		log.Fatal("Create table error: " + createUserTableErr.Error())
		return
	}

	log.Println("Check tables complete")
	closeErr := db.Close()
	if closeErr != nil {
		return
	}
}
func GetDbConnection() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", DbPath)
	if err != nil {
		log.Fatal("Connect DB error: " + err.Error())
	}
	return db
}
