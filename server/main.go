package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"Server/api"
	"Server/util"
	"Server/websocket"
)

/*
Program start

starts SQL;
creates WS and API;
starts listening and serving
*/
func main() {
	err := util.LoadConfig()
	if err != nil {
		util.Err(util.MAIN, err, true, "Error reading Configs")
		return
	}

	SQLInit()

	router := mux.NewRouter().StrictSlash(true)

	api.CreateAPI(router)
	websocket.Createwebsocket(router)

	util.Log("MAIN", "Started")

	// Blocking
	err = http.ListenAndServe(":"+fmt.Sprintf("%d", util.GetConfig().Port), router)
	util.Err(util.MAIN, err, true, "Listening Error")
}
