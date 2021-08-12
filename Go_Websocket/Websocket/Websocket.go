package websocket

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"Server/util"
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
			util.Log("WS", "error in reciving: ", err)
			fmt.Println()
			return
		}
		fmt.Println()
		util.Log("WS", c.RemoteAddr().String()+" -> "+string(raw))

		var recive map[string]interface{}
		err = json.Unmarshal(raw, &recive)

		if err != nil {
			util.Log("WS", "JSON decoding error:", err, " :", string(raw))
			continue
		}
		if len(recive) == 0 {
			util.Log("WS", "empty JSON")
			continue
		}
		Program_id, action := validateWSJSON(&recive)
		if action == "" {
			util.Log("WS", "invalid JSON WS request", recive)
			continue
		}
		util.Log("WS", "recived:", recive)

		var data interface{}
		var actionerr error

		switch action {
		case "Logs":
			actionerr = Checkadmin(&recive)
			if actionerr == nil {
				data, actionerr = Getlogs(Program_id)
			}
		case "Activity":
			data, actionerr = Getactivity(Program_id)
		case "Start":
			actionerr = Checkadmin(&recive)
			if actionerr == nil {
				data, actionerr = Start(Program_id)
			}
		case "Stop":
			actionerr = Checkadmin(&recive)
			if actionerr == nil {
				data, actionerr = Stop(Program_id)
			}
		}

		if actionerr != nil {
			util.Log("WS", "send err:", actionerr)
			msg, _ := json.Marshal(map[string]interface{}{"action": action, "error": actionerr.Error()})
			c.WriteMessage(1, msg)
		} else {
			msg, err := json.Marshal(map[string]interface{}{"action": action, "data": data})
			if err != nil {
				util.Log("WS", "send err:", err)
				msg, _ := json.Marshal(map[string]interface{}{"action": action, "error": err.Error()})
				c.WriteMessage(1, msg)
			}
			util.Log("WS", "send:", string(msg))
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
			util.Log("WS", err)
			util.Log("WS", "error while upgrading conncetion to websocket: ", r.RemoteAddr)
			return
		}
		util.Log("WS", "upgraded conncetion to websocket: ", r.RemoteAddr)
		go reciveWS(conn)
	})
}
