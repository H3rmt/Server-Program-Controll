package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"Server/api"
	"Server/util"
	"Server/websocket"
)

var Port = "18769"

/*
Program start

starts SQL;
creates WS and API;
starts listening and serving
*/
func main() {
	SQLInit()
	router := mux.NewRouter().StrictSlash(true)
	api.CreateAPI(router)
	websocket.Createwebsocket(router)
	util.Log("MAIN", "Started")
	err := http.ListenAndServe(":"+Port, router)
	util.Log("MAIN", "Err: ", err)
}
