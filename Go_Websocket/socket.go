package main

import (
	"encoding/json"
	"flag"
	"log"

	"net/http"

	"github.com/gorilla/websocket"

	"Go_Websocket/PRCommunication"
	"Go_Websocket/SQL"
)

var addr = flag.String("addr", ":18769", "")

var up = websocket.Upgrader{
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

/*
actions:

SQL:
	getlogs
	getactivity

Direct Programm communication (admin rights)
	start
	stop
	customaction
*/

func validateJSON(js *map[string]interface{}) bool {
	_, array_key_exists := (*js)["action"]
	return array_key_exists
}

func recive(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Print("error: ", err)
			}
			return
		}
		var recive map[string]interface{}
		log.Print(c.RemoteAddr().String() + "|" + string(message))
		json.Unmarshal(message, &recive)
		if err != nil {
			log.Print("error: ", err)
			continue
		} else if len(recive) == 0 {
			log.Print("empty JSON")
			continue
		} else if validateJSON(&recive) {
			log.Print("invalid JSON", recive)
			continue
		}
		log.Print(recive)
		switch recive["action"] {
		case "getlogs":
			log.Print("getLogs")
			SQL.Getlogs()
		case "getactivity":
			log.Print("getactivity")
			SQL.Getactivity()
		case "start":
			log.Print("start")
			PRCommunication.Start()
		case "stop":
			log.Print("stop")
			PRCommunication.Stop()
		case "customaction":
			log.Print("customaction")
			PRCommunication.Customaction()
		}
	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		log.Println(r.RemoteAddr)
		log.Println(conn)
		return
	}
	log.Print("upgraded conncetion to websocket: ", r.RemoteAddr)
	go recive(conn)
}

func main() {
	SQL.Init()
	flag.Parse()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serve(w, r)
	})
	log.Println("Started")
	err := http.ListenAndServe(*addr, nil)
	log.Print("ListenAndServe: ", err)
}
