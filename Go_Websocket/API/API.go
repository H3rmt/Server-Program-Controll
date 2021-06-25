package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	ws "Go_Websocket/WS"

	"github.com/gorilla/mux"
)

/*
checks if recived JSON has APIkey and ActivityRequest or LogRequest as keys
*/
func validateAPIJSON(js *map[string]interface{}) string {
	_, array_key_exists := (*js)["APIkey"]
	if array_key_exists {
		_, ActivityRequest_exists := (*js)["ActivityRequest"]
		_, LogRequest_exists := (*js)["LogRequest"]
		if ActivityRequest_exists && !LogRequest_exists {
			return "ActivityRequest"
		}
		if LogRequest_exists && !ActivityRequest_exists {
			return "LogRequest"
		}
	}
	return ""
}

/*
Error thrown/returned when no admin priv are present
*/
type InvalidAPIkeyerror struct{}

func (m *InvalidAPIkeyerror) Error() string {
	return "Invalid API key"
}

func reciveAPI(raw []byte, w *http.ResponseWriter) {
	fmt.Println()

	var recive map[string]interface{}
	err := json.Unmarshal(raw, &recive)

	if err != nil {
		log.Println("API|", "JSON decoding error: ", err)
		return
	}
	if len(recive) == 0 {
		log.Println("API|", "empty JSON")
		return
	}
	jsonvalidate := validateAPIJSON(&recive)
	if jsonvalidate == "" {
		log.Println("API|", "invalid JSON API request", recive)
		return
	}
	log.Println("API|", "recived: ", recive)

	var success = false

	switch jsonvalidate {
	case "ActivityRequest":
		success, err = ProcessActivityRequest(&recive)
	case "LogRequest":
		success, err = ProcessLogRequest(&recive)
	}

	if err != nil {
		log.Println("API|", "err:", err)
		APIerror, ok := err.(*InvalidAPIkeyerror)
		if ok {
			json.NewEncoder(*w).Encode(map[string]interface{}{"type": jsonvalidate, "error": APIerror})
		}
	} else {
		msg, _ := json.Marshal(map[string]interface{}{"type": jsonvalidate, "success": success})
		if err != nil {
			log.Println("API|", "err:", err)
			json.NewEncoder(*w).Encode(map[string]interface{}{"type": jsonvalidate, "error": "JSONerror"})
			return
		}
		(*w).Write(msg)
		log.Println("API|", "send:", string(msg))
	}
}

/*
Registers /api handle to mux.Router with json return and POST get
*/
func CreateAPI() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := ioutil.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
		go reciveAPI(raw, &w)
	}).Methods("POST")
	return router
}

/*
Api request:
{
	"APIkey":"gli23085uyljahlkhoql2emdga;fho8u3",
		"LogRequest":{
			"Date":"12.5.2012:13.52",
			"Number":123,
			"Message":"Test message",
			"Type":"Low",
		}
	/
		"ActivityRequest":{
			"Date":"12.5.2012:13.52",
			"Type":"Send",
		}
}
\

test:
curl -d {\"APIkey\":\"25253\", \"LogRequest\":\"12\"} http://localhost:18769/api
*/

/*
Struct to represent a Request asking to add a log in the Log table in DB
*/
type LogRequest struct {
	Date    string
	Number  int
	Message string
	Type    ws.Logtype
}

/*
Struct to represent a Request asking to add a activity in the Acitivity table in DB
*/
type ActivityRequest struct {
	Date string
	Type ws.Activitytype
}
