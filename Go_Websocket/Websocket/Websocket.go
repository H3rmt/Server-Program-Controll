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
func validateWSJSON(js *map[string]interface{}) (string, string) {
	Program_id, Program_id_key_exists := (*js)["Program_id"]
	if Program_id_key_exists {
		action := fmt.Sprintf("%v", (*js)["action"])
		if action == "Activity" || action == "Logs" || action == "Stop" || action == "Start" {
			return fmt.Sprintf("%v", Program_id), action
		}
	}
	return "", ""
}

/*
check if send shacode exists or equals stored sha code
*/
func checkadmin(js *map[string]interface{}) error {
	code, code_exists := (*js)["code"]
	if code_exists {
		valid := (code_exists && code == "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08") //test
		if valid {
			return nil
		}
	}
	return &Permissionerror{}
}

/*
Error thrown/returned when no admin priv are present
*/
type Permissionerror struct{}

func (m *Permissionerror) Error() string {
	return "no admin permissions"
}

/*
upgrades Connection and listens to incomming request

called as a go routine
*/
func reciveWS(c *websocket.Conn) {
	for {
		_, raw, err := c.ReadMessage()
		if err != nil {
			log.Println("WS|", "error in reciving: ", err)
			fmt.Println()
			return
		}
		fmt.Println()
		log.Println("WS|", c.RemoteAddr().String()+" -> "+string(raw))

		var recive map[string]interface{}
		err = json.Unmarshal(raw, &recive)

		if err != nil {
			log.Println("WS|", "JSON decoding error:", err, " :", string(raw))
			continue
		}
		if len(recive) == 0 {
			log.Println("WS|", "empty JSON")
			continue
		}
		Program_id, action := validateWSJSON(&recive)
		if action == "" {
			log.Println("WS|", "invalid JSON WS request", recive)
			continue
		}
		log.Println("WS|", "recived:", recive)

		var data interface{}
		var actionerr error

		switch action {
		case "Logs":
			data, actionerr = Getlogs(Program_id)
		case "Activity":
			data, actionerr = Getactivity(Program_id)
		case "Start":
			actionerr = checkadmin(&recive)
			if actionerr == nil {
				data, actionerr = Start(Program_id)
			}
		case "Stop":
			actionerr = checkadmin(&recive)
			if actionerr == nil {
				data, actionerr = Stop(Program_id)
			}
		}

		if actionerr != nil {
			log.Println("WS|", "send err:", actionerr)
			msg, _ := json.Marshal(map[string]interface{}{"action": action, "error": actionerr.Error()})
			c.WriteMessage(1, msg)
		} else {
			msg, err := json.Marshal(map[string]interface{}{"action": action, "data": data})
			if err != nil {
				log.Println("WS|", "send err:", err)
				msg, _ := json.Marshal(map[string]interface{}{"action": action, "error": err.Error()})
				c.WriteMessage(1, msg)
			}
			log.Println("WS|", "send:", string(msg))
			c.WriteMessage(1, msg)
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
