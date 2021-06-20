package main

import (
	"encoding/json"
	"fmt"
	"log"

	"flag"
	"net/http"

	"github.com/gorilla/websocket"

	"Go_Websocket/ExternalCommunication"
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
	return !array_key_exists
}

func recive(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		fmt.Println()
		if err != nil {
			log.Print("error in reciving: ", err)
			fmt.Println()
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
		log.Print("recived: ", recive)

		var returnval interface{}

		switch recive["action"] {
		case "getlogs":
			returnval, err = SQL.Getlogs(&recive)
		case "getactivity":
			returnval, err = SQL.Getactivity(&recive)
		case "start":
			returnval, err = ExternalCommunication.Start(&recive)
		case "stop":
			returnval, err = ExternalCommunication.Stop(&recive)
		case "customaction":
			returnval, err = ExternalCommunication.Customaction(&recive)
		}
		if err != nil {
			log.Println(err)
			_, ok := err.(*ExternalCommunication.Permissionerror)
			if ok {
				rec, _ := json.Marshal(map[string]interface{}{"action": recive["action"], "error": "Permissionerror"})
				c.WriteMessage(1, rec)
			}
		} else {
			rec, err := json.Marshal(map[string]interface{}{"action": recive["action"], "data": returnval})
			if err != nil {
				log.Println(err)
				rec, _ := json.Marshal(map[string]interface{}{"action": recive["action"], "error": "JSONerror"})
				c.WriteMessage(1, rec)
			} else {
				log.Println("sending: ", string(rec))
				c.WriteMessage(1, rec)
			}
		}
	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		log.Println("error while upgrading conncetion to websocket: ", r.RemoteAddr)
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
	log.Print("Err: ", err)
}
