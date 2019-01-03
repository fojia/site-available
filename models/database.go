package models

import "database/sql"

var DB *sql.DB
var Err error
//Connection to db
func InitDB() {
	DB, Err = sql.Open("sqlite3", "./database.sqlite")
	CheckErr(Err)
}

// Summarize errors for reduce code
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}

}
