package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/gorilla/mux"

	// create struct from map out of JSON
	"github.com/mitchellh/mapstructure"
)

/*
checks if recived JSON has APIkey and ActivityRequest or LogRequest as keys
returns the API key and REquesttype
*/
func validateAPIJSON(js *map[string]interface{}) (string, string) {
	APIKey, array_key_exists := (*js)["APIkey"]
	if array_key_exists {
		_, ActivityRequest_exists := (*js)["ActivityRequest"]
		_, LogRequest_exists := (*js)["LogRequest"]
		if ActivityRequest_exists && !LogRequest_exists {
			return APIKey.(string), "ActivityRequest"
		}
		if LogRequest_exists && !ActivityRequest_exists {
			return APIKey.(string), "LogRequest"
		}
	}
	return APIKey.(string), ""
}

/*
Error thrown/returned when no admin priviges are present
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
	APIKey, request := validateAPIJSON(&recive)
	if request == "" {
		log.Println("API|", "invalid JSON API request", recive)
		return nil
	}
	log.Println("API|", "recived: ", recive)

	switch request {
	case "ActivityRequest":
		var activityrequest ActivityRequest
		mapstructure.Decode(recive["ActivityRequest"], &activityrequest)
		err = ProcessActivityRequest(APIKey, &activityrequest)
	case "LogRequest":
		var logrequest LogRequest
		mapstructure.Decode(recive["LogRequest"], &logrequest)
		err = ProcessLogRequest(APIKey, &logrequest)
	}

	if err != nil {
		log.Println("API|", "err:", err)
		APIerror, ok := err.(*InvalidAPIkeyerror)
		if ok {
			msg, _ := json.Marshal(map[string]interface{}{"type": request, "error": APIerror})
			return msg
		}
		return nil
	} else {
		msg, _ := json.Marshal(map[string]interface{}{"type": request, "success": true})
		if err != nil {
			log.Println("API|", "err:", err)
			msg, _ := json.Marshal(map[string]interface{}{"type": request, "error": "JSONerror"})
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

		msg := reciveAPI(&raw)

		if msg == nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("API|", "send: bad request")
			return
		}
		_, err := w.Write(msg)
		if err != nil {
			log.Println("API|", "err in sending:", err)
		}

		log.Println("API|", "send:", string(msg))
	}).Methods("POST")
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
curl -d {\"APIkey\":\"25253\", \"LogRequest\":1} http://localhost:18769/api
*/
