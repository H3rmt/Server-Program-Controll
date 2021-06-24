package main

import (
	api "Go_Websocket/API"
	ws "Go_Websocket/WS"
	"log"
	"net/http"
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
	router := api.CreateAPI()
	ws.Createwebsocket(router)
	log.Println("MAIN|", "Started")
	err := http.ListenAndServe(":"+Port, router)
	log.Println("MAIN|", "Err: ", err)
}
