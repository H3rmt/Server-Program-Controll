package SQL

import (
	"context"
	"database/sql"

	f "fmt"
	"log"
)

var db *sql.DB

var server = "127.0.0.1"
var port = 3306
var user = "Go"
var password = "e73EG6dP2f8F2dAx"
var database = "programs"

// Exported
func Init() {
	db, err := sql.Open("sqlserver", f.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database))
	if err != nil {
		log.Fatal("Error creating connection: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Connected!\n", db)
}

// Exported
func Getlogs() {
	log.Print("Connected!\n", db)
}

// Exported
func Getactivity() {

}
