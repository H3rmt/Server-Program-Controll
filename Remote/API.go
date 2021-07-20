package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

var Port = "18770"

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
	APIKey, Api_key_exists := (*js)["APIkey"]
	if Api_key_exists {
		return APIKey.(string)
	}
	return ""
}

/*
finds the corresponding program from list of programs
*/
func getProgramm_IDfromAPIKey(APIKey string) (*Program, error) {
	for i := 0; i < len(programs); i++ {
		if programs[i].APIKey == APIKey {
			log.Println("Program found:", APIKey, programs[i].Program, programs[i].Arguments)
			return &programs[i], nil
		}
	}
	return &Program{}, &InvalidAPIkeyerror{}
}

/*
Error returned when APIkey was invalid
*/
type InvalidAPIkeyerror struct{}

func (m *InvalidAPIkeyerror) Error() string {
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

	var commandRequest ActualCommandRequest
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
Process request to execute command in Program
*/
func ProcessCommandRequest(program *Program, request *ActualCommandRequest) (err error) {
	switch request.Message {
	case "Start":
		err = program.Start()
	case "Stop":
		err = program.Stop()
	default:
		err = fmt.Errorf("unsuported Comand")
	}
	return
}

/*
Struct to represent a Request send to actual program asking to execute command
*/
type ActualCommandRequest struct {
	Message string
	APIkey  string
}

/*
Struct to represent a Request asking to add a log in the Log table in DB
*/
type LogRequest struct {
	Date    string
	Number  int
	Message string
	Type    Logtype
}

type Logtype string

const (
	Low       Logtype = "Low"
	Normal    Logtype = "Normal"
	Important Logtype = "Important"
	Error     Logtype = "Error"
)

/*
Struct to represent a Request asking to add a activity in the Acitivity table in DB
*/
type ActivityRequest struct {
	Date string
	Type Activitytype
}

type Activitytype string

const (
	Send              Activitytype = "Send"
	Recive            Activitytype = "Recive"
	Process           Activitytype = "Process"
	Backgroundprocess Activitytype = "Backgroundprocess"
)

// curl -d {\"APIkey\":\"4362fds357rd32q1f37y35e6ytefws\",\"Message\":\"Stop\"} http://localhost:18770/api

/*
	log.Println("Performing Http Post...")
	req := map[string]interface{}{"APIkey": "5rrtg3u564uiqr43fadf", "Log": LogRequest{Date: "2021-06-17 22:26:43", Number: 12, Message: "hitext", Type: Important}}
	jsonReq, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:18769/api", "application/json;", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Println(bodyString)
*/
