package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"net/http"

	"github.com/gorilla/mux"

	// create struct from map out of JSON
	"github.com/mitchellh/mapstructure"
)

/*
checks if recived JSON has APIkey and (Register or Activity or Log or Customaction) as keys
returns the API key and Requesttype
*/
func validateAPIJSON(js *map[string]interface{}) (string, string) {
	APIKey, Api_key_exists := (*js)["APIkey"]
	if Api_key_exists {
		_, Register_exists := (*js)["Register"]
		_, Activity_exists := (*js)["Activity"]
		_, Log_exists := (*js)["Log"]
		_, Action_exists := (*js)["Action"]
		if Register_exists && !Activity_exists && !Log_exists && !Action_exists {
			return APIKey.(string), "Register"
		}
		if Activity_exists && !Register_exists && !Log_exists && !Action_exists {
			return APIKey.(string), "Activity"
		}
		if Log_exists && !Register_exists && !Activity_exists && !Action_exists {
			return APIKey.(string), "Log"
		}
		if Action_exists && !Register_exists && !Activity_exists && !Log_exists {
			return APIKey.(string), "Action"
		}
	}
	return "", ""
}

/*
Error returned when APIkey was invalid
*/
type InvalidAPIkeyerror struct{}

func (m *InvalidAPIkeyerror) Error() string {
	return "Invalid API key"
}

/*
Error returned when Request values were invalid
*/
type InvalidRequesterror struct{}

func (m *InvalidRequesterror) Error() string {
	return "Invalid Request values"
}

/*
called when Connection send data;
gets byte array out of JSON
returns byte array out of JSON to write
*/
func reciveAPI(raw *[]byte, addr string) []byte {
	fmt.Println()

	var recive map[string]interface{}
	err := json.Unmarshal(*raw, &recive)

	if err != nil {
		log.Println("API|", "JSON decoding error: ", err, " :", string(*raw))
		return nil
	}
	if len(recive) == 0 {
		log.Println("API|", "empty JSON")
		return nil
	}
	APIKey, request := validateAPIJSON(&recive)
	if request == "" {
		log.Println("API|", "invalid JSON API request", recive)
		return nil
	}
	log.Println("API|", "recived: ", recive)

	ProgammID, err := getProgramm_IDfromAPIKey(APIKey)
	if err != nil {
		log.Println("API|", "APIKeyerr:", err)
		msg, _ := json.Marshal(map[string]interface{}{"type": request, "error": err.Error()})
		return msg
	}

	switch request {
	case "Register":
		log.Println(addr)
		addr := strings.TrimRight(addr, "0123456789:")
		log.Println(addr)
		err = ProcessRegisterRequest(ProgammID, addr)
	case "Activity":
		activityrequest := ActivityRequest{Date: "-1", Type: "-1"}
		mapstructure.Decode(recive["ActivityRequest"], &activityrequest)
		if activityrequest.Date != "-1" && activityrequest.Type != "-1" {
			err = ProcessActivityRequest(ProgammID, &activityrequest)
		} else {
			err = &InvalidRequesterror{}
		}
	case "Log":
		logrequest := LogRequest{Date: "-1", Number: -1, Message: "-1", Type: "-1"}
		mapstructure.Decode(recive["LogRequest"], &logrequest)
		if logrequest.Date != "-1" && logrequest.Number != -1 && logrequest.Message != "-1" && logrequest.Type != "-1" {
			err = ProcessLogRequest(ProgammID, &logrequest)
		} else {
			err = &InvalidRequesterror{}
		}
	case "Action":
		commandrequest := CommandRequest{Message: "-1"}
		mapstructure.Decode(recive["CommandRequest"], &commandrequest)
		if commandrequest.Message != "-1" {
			err = ProcessCommandRequest(ProgammID, &commandrequest, APIKey)
		} else {
			err = &InvalidRequesterror{}
		}
	}

	if err != nil {
		log.Println("API|", "err:", err)
		msg, _ := json.Marshal(map[string]interface{}{"type": request, "error": err.Error()})
		return msg
	} else {
		msg, _ := json.Marshal(map[string]interface{}{"type": request, "success": true})
		if err != nil {
			log.Println("API|", "err:", err)
			msg, _ := json.Marshal(map[string]interface{}{"type": request, "error": err.Error()})
			return msg
		}
		return msg
	}
}

/*
Registers /api handle to mux.Router with json return and POST get
*/
func CreateAPI(rout *mux.Router) {
	rout.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := ioutil.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")

		msg := reciveAPI(&raw, r.RemoteAddr)

		if msg == nil {
			w.WriteHeader(http.StatusBadRequest)
			msg, _ = json.Marshal(map[string]interface{}{"error": "bad request"})
		}

		_, err := w.Write(msg)
		if err != nil {
			log.Println("API|", "err in sending:", err)
		} else {
			log.Println("API|", "send:", string(msg))
		}
	}).Methods("POST")
}

/*
Api request:
{
	"APIkey":"gli23085uyljahlkhoql2emdga;fho8u3",
		"Log":{
			"Date":"12.5.2012:13.52",
			"Number":123,
			"Message":"Test message",
			"Type":"Low",
		}
	/
		"Activity":{
			"Date":"12.5.2012:13.52",
			"Type":"Send",
		}
}
\

test:
curl -d {\"APIkey\":\"25253\",\"Log\":1} http://localhost:18769/api
*/
