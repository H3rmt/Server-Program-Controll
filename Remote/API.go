package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

var Port = "18770"
var RemotePort = "18769"

/*
start API to listen to CommandRequests
*/
func CreateAPI(rout *mux.Router) {
	rout.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := ioutil.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")

		msg := reciveAPI(&raw)

		if msg == nil {
			w.WriteHeader(http.StatusBadRequest)
			msg, _ = json.Marshal(map[string]interface{}{"error": "bad request"})
		}

		_, err := w.Write(msg)
		if err != nil {
			log.Println("err in sending:", err)
		} else {
			log.Println("send:", string(msg))
		}
	}).Methods("POST")
}

func validateAPIJSON(js *map[string]interface{}) string {
	APIKey, Api_key_exists := (*js)["APIKey"]
	if Api_key_exists {
		return APIKey.(string)
	}
	return ""
}

/*
Error returned when APIKey was invalid
*/
type InvalidAPIKeyerror struct{}

func (m *InvalidAPIKeyerror) Error() string {
	return "Invalid API key"
}

/*
called when Connection send data;
gets byte array out of JSON
returns byte array out of JSON to write
*/
func reciveAPI(raw *[]byte) []byte {
	fmt.Println()

	var recive map[string]interface{}
	err := json.Unmarshal(*raw, &recive)

	if err != nil {
		log.Println("API|", "JSON decoding error: ", err)
		return nil
	}
	if len(recive) == 0 {
		log.Println("API|", "empty JSON")
		return nil
	}
	APIKey := validateAPIJSON(&recive)
	if APIKey == "" {
		log.Println("API|", "invalid JSON API request", recive)
		return nil
	}
	log.Println("API|", "recived: ", recive)

	Program, err := getProgramm_IDfromAPIKey(APIKey)
	if err != nil {
		log.Println("API|", "err:", err)
		msg, _ := json.Marshal(map[string]interface{}{"error": err.Error()})
		return msg
	}

	var commandRequest CommandRequest
	mapstructure.Decode(recive, &commandRequest)
	err = ProcessCommandRequest(Program, &commandRequest)

	if err != nil {
		log.Println("API|", "err:", err)
		msg, _ := json.Marshal(map[string]interface{}{"error": err.Error()})
		return msg
	} else {
		msg, _ := json.Marshal(map[string]interface{}{"success": true})
		if err != nil {
			log.Println("API|", "err:", err)
			msg, _ := json.Marshal(map[string]interface{}{"error": err.Error()})
			return msg
		}
		return msg
	}
}

/*
called to register a program with an APIKey to remote server
*/
func register(remote string, APIKey string) error {
	log.Println("Registering Program:", APIKey, " on", remote)
	req := map[string]interface{}{"APIKey": APIKey, "Register": true}
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("http://%s:%s/api", remote, RemotePort), "application/json;", bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println("recived:", string(bodyBytes))

	answer := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		return err
	}

	if answer["error"] != nil {
		return fmt.Errorf(answer["error"].(string))
	}
	return nil
}
