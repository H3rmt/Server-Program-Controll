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
	PDB, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", util.GetConfig().User, util.GetConfig().Password, util.GetConfig().DbPort, util.GetConfig().Database))
	if err != nil {
		util.Err("SQL", err, true, "Error creating connection")
		panic(err)
	}
	ADB, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", util.GetConfig().User, util.GetConfig().Password, util.GetConfig().DbPort, util.GetConfig().Database2))
	if err != nil {
		util.Err("SQL", err, true, "Error creating connection")
		panic(err)
	}

	err = PDB.Ping()
	if err != nil {
		util.Err("SQL", err, true, "Error creating PDB connection")
		panic(err)
	}

	err = ADB.Ping()
	if err != nil {
		util.Err("SQL", err, true, "Error creating ADB connection")
		panic(err)
	}
	util.Log("SQL", "Connected!")

	websocket.SetPDB(PDB)
	websocket.SetADB(ADB)
	api.SetPDB(PDB)
}
