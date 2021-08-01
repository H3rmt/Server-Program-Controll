package main

import (
	"fmt"

	"context"
	"database/sql"

	// used at sql.Open(->"mysql"<-, fmt.Sprintf
	_ "github.com/go-sql-driver/mysql"

	"Server/api"
	"Server/util"
	"Server/websocket"
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
		util.Log("SQL", "Error creating connection: ", err.Error())
		panic(err)
	}
	ctx := context.Background()
	err = DB.PingContext(ctx)
	if err != nil {
		util.Log("SQL", err.Error())
		panic(err)
	}
	util.Log("SQL", "Connected!")
	websocket.SetDB(DB)
	api.SetDB(DB)
}
