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
func validateWSJSON(js *map[string]any) (string, string) {
	id, ProgramIdKeyExists := (*js)["id"]
	if ProgramIdKeyExists {
		action := fmt.Sprintf("%v", (*js)["action"])
		if action == "Activity" || action == "Logs" || action == "Stop" || action == "Start" {
			return fmt.Sprintf("%v", id), action
		}
	}
	return "", ""
}

/*
Permissionerror thrown/returned when insufficient or no permissions are present
*/
type Permissionerror struct{}

func (m *Permissionerror) Error() string {
	return "missing permissions"
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

		var recive map[string]any
		err = json.Unmarshal(raw, &recive)

		if err != nil {
			util.Log("WS", "JSON decoding error:", err, " :", string(raw))
			continue
		}
		if len(recive) == 0 {
			util.Log("WS", "empty JSON")
			continue
		}
		id, action := validateWSJSON(&recive)
		if action == "" {
			util.Log("WS", "invalid JSON WS request", recive)
			continue
		}
		util.Log("WS", "recived:", recive)

		var data any
		var actionErr error
		actionErr = CheckPermission(&recive)
		if actionErr == nil {
			switch action {
			case "Logs":
				data, actionErr = Getlogs(id)
			case "Activity":
				data, actionErr = Getactivity(id)
			case "Start":
				data, actionErr = Start(id)
			case "Stop":
				data, actionErr = Stop(id)
			}
		}

		if actionErr != nil {
			util.Log("WS", "send err:", actionErr)
			msg, _ := json.Marshal(map[string]any{"action": action, "error": actionErr.Error()})
			c.WriteMessage(1, msg)
		} else {
			msg, err := json.Marshal(map[string]any{"action": action, "data": data})
			if err != nil {
				util.Log("WS", "send err:", err)
				msg, _ := json.Marshal(map[string]any{"action": action, "error": err.Error()})
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
			util.Log("WS", "error while upgrading connection to websocket: ", r.RemoteAddr)
			return
		}
		// util.Log("WS", "upgraded connection to websocket: ", r.RemoteAddr)
		go reciveWS(conn)
	})
}
