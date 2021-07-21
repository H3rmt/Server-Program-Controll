package main

import (
	api "Go_Websocket/API"
	ws "Go_Websocket/Websocket"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	ws.Createwebsocket(router)
	log.Println("MAIN|", "Started")
	err := http.ListenAndServe(":"+Port, router)
	log.Println("MAIN|", "Err: ", err)
}
