package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var Port = "18769"

var up = websocket.Upgrader{
	// allow connections from outside
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

/*
checkt if recived JSON has action als key
*/
func validateJSON(js *map[string]interface{}) bool {
	_, array_key_exists := (*js)["action"]
	return !array_key_exists
}

/*
upgrades Connection and listens to incomming request

called as a go routine
*/
func recive(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		fmt.Println()
		if err != nil {
			log.Println("WS|", "error in reciving: ", err)
			fmt.Println()
			return
		}

		var recive map[string]interface{}
		log.Println("WS|", c.RemoteAddr().String()+" | "+string(message))
		err = json.Unmarshal(message, &recive)
		if err != nil {
			log.Println("WS|", "JSON decoding error: ", err)
			continue
		} else if len(recive) == 0 {
			log.Println("WS|", "empty JSON")
			continue
		} else if validateJSON(&recive) {
			log.Println("WS|", "invalid JSON request", recive)
			continue
		}
		log.Println("WS|", "recived: ", recive)

		var returnval interface{}

		switch recive["action"] {
		case "getlogs":
			returnval, err = Getlogs(&recive)
		case "getactivity":
			returnval, err = Getactivity(&recive)
		case "start":
			returnval, err = Start(&recive)
		case "stop":
			returnval, err = Stop(&recive)
		case "customaction":
			returnval, err = Customaction(&recive)
		}
		if err != nil {
			log.Println("WS|", err)
			_, ok := err.(*Permissionerror)
			if ok {
				rec, _ := json.Marshal(map[string]interface{}{"action": recive["action"], "error": "Permissionerror"})
				c.WriteMessage(1, rec)
			}
		} else {
			rec, err := json.Marshal(map[string]interface{}{"action": recive["action"], "data": returnval})
			if err != nil {
				log.Println("WS|", err)
				rec, _ := json.Marshal(map[string]interface{}{"action": recive["action"], "error": "JSONerror"})
				c.WriteMessage(1, rec)
			} else {
				log.Println("WS|", "sending: ", string(rec))
				c.WriteMessage(1, rec)
			}
		}
	}
}

/*
Registers /ws handle to http to create websocket and send and recive JSON
*/
func createwebsocket(r *mux.Router) {
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := up.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WS|", err)
			log.Println("WS|", "error while upgrading conncetion to websocket: ", r.RemoteAddr)
			return
		}
		log.Println("WS|", "upgraded conncetion to websocket: ", r.RemoteAddr)
		go recive(conn)
	})
}

/*
Program start

starts SQL;
creates WS and API;
starts listening and serving
*/
func main() {
	SQLInit()
	router := createAPI()
	createwebsocket(router)
	log.Println("MAIN|", "Started")
	err := http.ListenAndServe(":"+Port, router)
	log.Println("MAIN|", "Err: ", err)
}
