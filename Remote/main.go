package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"Remote/api"
	"Remote/util"
)

var remoteIP string

/*
main method with arguments = programfile APIKey programfile APIKey ...
*/
func main() {
	data, err := ioutil.ReadFile("programs.json")
	if err != nil {
		util.Log("MAIN", "programs.json read error: ", err)
		panic(err)
	}

	var fileload map[string]interface{}
	json.Unmarshal(data, &fileload)

	for _, v := range fileload["Programs"].([]interface{}) {
		p := v.(map[string]interface{})
		var name = p["Name"].(string)
		var program = p["Program"].(string)
		var key = p["APIKey"].(string)
		var args = make([]string, len(p["Arguments"].([]interface{})))
		for count, v := range p["Arguments"].([]interface{}) {
			args[count] = v.(string)
		}
		if program != "" && key != "" {
			api.Programs = append(api.Programs, api.Program{Name: name, Program: program, APIKey: key, Arguments: args})
		} else {
			util.Log("MAIN", "Invalid Program:", program, key, args)
		}
	}
	util.Log("MAIN", "loaded Programs:")
	for _, v := range api.Programs {
		fmt.Printf("%s %s Key:%s \n", v.Program, v.Arguments, v.APIKey)
	}
	fmt.Println()

	router := mux.NewRouter().StrictSlash(true)
	api.CreateAPI(router)
	util.Log("MAIN", "Started API")

	remoteIP = fileload["Remote IP"].(string)
	api.SetRemoteIP(remoteIP)
	for _, v := range api.Programs {
		err = api.Register(remoteIP, v.APIKey)
		if err != nil {
			panic(err)
		}
	}

	// Blocking
	err = http.ListenAndServe(":"+api.Port, router)
	util.Log("MAIN", "Listening Error", err)
}
