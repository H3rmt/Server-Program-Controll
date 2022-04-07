package main

import (
	"fmt"
	// used at sql.Open(->"mysql"<-, fmt.Sprintf
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"Server/api"
	"Server/util"
	"Server/websocket"
)

/*
Create and open SQL Connection
*/
func SQLInit() {
	DB, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", util.GetConfig().User, util.GetConfig().Password, util.GetConfig().DbPort, util.GetConfig().Database))
	if err != nil {
		util.Err("SQL", err, true, "Error creating connection")
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		util.Err("SQL", err, true, "Error creating connection")
		panic(err)
	}
	util.Log("SQL", "Connected!")

	websocket.SetDB(DB)
	api.SetDB(DB)
}
