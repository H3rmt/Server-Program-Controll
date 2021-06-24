package ws

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var up = websocket.Upgrader{
	// allow connections from outside
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

/*
checks if recived JSON has action as key
*/
func validateWSJSON(js *map[string]interface{}) bool {
	_, array_key_exists := (*js)["action"]
	return array_key_exists
}

/*
upgrades Connection and listens to incomming request

called as a go routine
*/
func reciveWS(c *websocket.Conn) {
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
		}
		if len(recive) == 0 {
			log.Println("WS|", "empty JSON")
			continue
		}
		if !validateWSJSON(&recive) {
			log.Println("WS|", "invalid JSON WS request", recive)
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
		default:
			err = fmt.Errorf("unknown Action")
		}

		if err != nil {
			log.Println("WS|", "err:", err)
			WSerror, ok := err.(*Permissionerror)
			if ok {
				msg, _ := json.Marshal(map[string]interface{}{"action": recive["action"], "error": WSerror})
				c.WriteMessage(1, msg)
			}
		} else {
			msg, err := json.Marshal(map[string]interface{}{"action": recive["action"], "data": returnval})
			if err != nil {
				msg, _ := json.Marshal(map[string]interface{}{"action": recive["action"], "error": "JSONerror"})
				c.WriteMessage(1, msg)
				return
			}
			c.WriteMessage(1, msg)
			log.Println("WS|", "send:", string(msg))
		}
	}
}

/*
Registers /ws handle to http to create websocket and send and recive JSON
*/
func Createwebsocket(rout *mux.Router) {
	rout.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := up.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WS|", err)
			log.Println("WS|", "error while upgrading conncetion to websocket: ", r.RemoteAddr)
			return
		}
		log.Println("WS|", "upgraded conncetion to websocket: ", r.RemoteAddr)
		go reciveWS(conn)
	})
}
