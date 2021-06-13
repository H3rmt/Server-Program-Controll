package main

import (
	"encoding/json"
	"flag"
	"log"

	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":18769", "")

var up = websocket.Upgrader{
	CheckOrigin: func(*http.Request) bool {
		return true
	},
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
		}
		log.Print(recive)
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
	flag.Parse()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serve(w, r)
	})
	log.Println("Started")
	err := http.ListenAndServe(*addr, nil)
	log.Print("ListenAndServe: ", err)
}
