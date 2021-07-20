package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// list of all programs
var programs []Program

/*
main method with arguments = programfile APIKey programfile APIKey ...
*/
func main() {
	// Create programs from args

	data, err := ioutil.ReadFile("programs.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var fileload map[string]interface{}
	json.Unmarshal(data, &fileload)

	for _, v := range fileload["Programs"].([]interface{}) {
		p := v.(map[string]interface{})
		var program = p["Program"].(string)
		var key = p["APIKey"].(string)
		var args = make([]string, len(p["Arguments"].([]interface{})))
		for count, v := range p["Arguments"].([]interface{}) {
			args[count] = v.(string)
		}
		if program != "" && key != "" {
			programs = append(programs, Program{Program: program, APIKey: key, Arguments: args})
		} else {
			log.Println("Invalid Program:", program, key, args)
		}
	}
	log.Println("loaded Programs:")
	for _, v := range programs {
		fmt.Printf("%s %s Key:%s \n", v.Program, v.Arguments, v.APIKey)
	}

	router := mux.NewRouter().StrictSlash(true)
	CreateAPI(router)
	log.Println("Started API")

	err = http.ListenAndServe(":"+Port, router)
	log.Println("Err: ", err)
}
