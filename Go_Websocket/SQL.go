package main

import (
	"context"
	"database/sql"
	"fmt"
	
	// used at sql.Open(->"mysql"<-, fmt.Sprintf
	_ "github.com/go-sql-driver/mysql"
	
	"Server/api"
	"Server/util"
	"Server/websocket"
)

/*
Create and open SQL Connection
*/
func SQLInit() {
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", util.GetConfig().User, util.GetConfig().Password, util.GetConfig().Database))
	if err != nil {
		util.Err("SQL", err, true, "Error creating connection")
	}
	ctx := context.Background()
	err = DB.PingContext(ctx)
	if err != nil {
		util.Err("SQL", err, true, "Error creating connection")
	}
	util.Log("SQL", "Connected!")

	websocket.SetDB(DB)
	api.SetDB(DB)
}
