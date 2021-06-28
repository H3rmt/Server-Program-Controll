package main

import (
	api "Go_Websocket/API"
	ws "Go_Websocket/WS"
	"fmt"
	"log"

	"context"
	"database/sql"

	// used at sql.Open(->"mysql"<-, fmt.Sprintf
	_ "github.com/go-sql-driver/mysql"
)

var user = "Go"
var password = "e73EG6dP2f8F2dAx"
var database = "programs"

/*
Create and open SQL Connection
*/
func SQLInit() {
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, database))
	if err != nil {
		log.Println("SQL|", "Error creating connection: ", err.Error())
	}
	ctx := context.Background()
	err = DB.PingContext(ctx)
	if err != nil {
		log.Println("SQL|", err.Error())
	}
	log.Println("SQL|", "Connected!")
	ws.SetDB(DB)
	api.SetDB(DB)
}
