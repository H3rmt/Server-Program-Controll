package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/mitchellh/mapstructure"

	"Server/util"
)

/*
checks if recived JSON has APIkey and (Register or Activity or Log or Action) as key
returns the API key and Requesttype
*/
func validateAPIJSON(js *map[string]any) (string, action string) {
	APIKey, ApiKeyExists := (*js)["APIKey"]
	if ApiKeyExists {
		_, RegisterExists := (*js)["Register"]
		_, ActivityExists := (*js)["Activity"]
		_, LogExists := (*js)["Log"]
		_, StatechangeExists := (*js)["StateChange"]
		if RegisterExists && !ActivityExists && !LogExists && !StatechangeExists {
			return fmt.Sprintf("%v", APIKey), "Register"
		}
		if ActivityExists && !RegisterExists && !LogExists && !StatechangeExists {
			return fmt.Sprintf("%v", APIKey), "Activity"
		}
		if LogExists && !RegisterExists && !ActivityExists && !StatechangeExists {
			return fmt.Sprintf("%v", APIKey), "Log"
		}
		if StatechangeExists && !RegisterExists && !ActivityExists && !LogExists {
			return fmt.Sprintf("%v", APIKey), "StateChange"
		}
	}
	return
}

/*
Error returned when Request values were invalid
*/
type InvalidRequestError struct{}

func (m *InvalidRequestError) Error() string {
	return "Invalid Request values"
}

/*
called when Connection send data;
gets byte array out of JSON
returns byte array out of JSON to write
*/
func reciveAPI(raw *[]byte, addr string) []byte {
	fmt.Println()

	var recive map[string]any
	err := json.Unmarshal(*raw, &recive)

	if err != nil {
		util.Log("API", "JSON decoding error: ", err, " :", string(*raw))
		return nil
	}
	if len(recive) == 0 {
		util.Log("API", "empty JSON")
		return nil
	}
	APIKey, request := validateAPIJSON(&recive)
	if request == "" {
		util.Log("API", "invalid JSON API request", recive)
		return nil
	}
	util.Log("API", "recived:", recive)

	ProgammID, err := getProgram_IDfromAPIKey(APIKey)
	if err != nil {
		util.Log("API", "APIKeyerr:", err)
		msg, _ := json.Marshal(map[string]any{"type": request, "error": err.Error()})
		return msg
	}

	switch request {
	case "Register":
		splitaddr := strings.SplitN(addr, ":", -1)
		addr := strings.Replace(addr, ":"+splitaddr[len(splitaddr)-1], "", 1) // remove client port
		err = ProcessRegisterRequest(ProgammID, addr, uint16(int(recive["Port"].(float64))))
	case "Activity":
		activityrequest := ActivityRequest{Date: "-1", Type: "-1"}
		mapstructure.Decode(recive["Activity"], &activityrequest)
		if activityrequest.Date != "-1" && activityrequest.Type != "-1" {
			err = ProcessActivityRequest(ProgammID, &activityrequest)
		} else {
			err = &InvalidRequestError{}
		}
	case "Log":
		logrequest := LogRequest{Date: "-1", Message: "-1", Type: "-1"}
		mapstructure.Decode(recive["Log"], &logrequest)
		if logrequest.Date != "-1" && logrequest.Message != "-1" && logrequest.Type != "-1" {
			err = ProcessLogRequest(ProgammID, &logrequest)
		} else {
			err = &InvalidRequestError{}
		}
	case "StateChange":
		statechangerequest := StateChangeRequest{Date: "-1"}
		mapstructure.Decode(recive["StateChange"], &statechangerequest)
		if statechangerequest.Date != "-1" {
			err = ProcessStateChangeRequest(ProgammID, &statechangerequest)
		} else {
			err = &InvalidRequestError{}
		}

	}

	if err != nil {
		util.Log("API", "send err:", err)
		msg, _ := json.Marshal(map[string]any{"type": request, "error": err.Error()})
		return msg
	} else {
		msg, _ := json.Marshal(map[string]any{"type": request, "success": true})
		if err != nil {
			util.Log("API", "send err:", err)
			msg, _ := json.Marshal(map[string]any{"type": request, "error": err.Error()})
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
			msg, _ = json.Marshal(map[string]any{"error": "bad request"})
		}

		_, err := w.Write(msg)
		if err != nil {
			util.Log("API", "err in sending:", err)
		} else {
			util.Log("API", "send:", string(msg))
		}
	}).Methods("POST")
}
