package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

/*
Registers /api handle to mux.Router with json return and POST get
*/
func createAPI() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println()

		raw, _ := ioutil.ReadAll(r.Body)

		var recive map[string]interface{}
		err := json.Unmarshal(raw, &recive)

		log.Println("API|", "recived: ", string(raw), recive, err)
		fmt.Println()

		json.NewEncoder(w).Encode(map[string]interface{}{"id send: ": recive["id"]})
	}).Methods("POST")
	return router
}

/*
Api request:
{
	"Apikey":"gli23085uyljahlkhoql2emdga;fho8u3",
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

*/

/*
Struct to represent a Request asking to add a log in the Log table in DB
*/
type LogRequest struct {
	Date    string
	Number  int
	Message string
	Type    Logtype
}

/*
Struct to represent a Request asking to add a activity in the Acitivity table in DB
*/
type ActivityRequest struct {
	Date string
	Type Activitytype
}
