package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func createAPI() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		raw, _ := ioutil.ReadAll(r.Body)

		var recive map[string]interface{}
		err := json.Unmarshal(raw, &recive)

		log.Println(string(raw), recive, err)

		json.NewEncoder(w).Encode(map[string]interface{}{"id send: ": recive["id"]})
	}).Methods("POST")
	return router
}
