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

/*
main method with arguments = programfile APIKey programfile APIKey ...
*/
func main() {
	err := util.LoadConfig()
	if err != nil {
		util.Err(util.MAIN, err, true, "Error reading Configs")
		return
	}

	data, err := ioutil.ReadFile("programs.json")
	if err != nil {
		util.Err(util.MAIN, err, true, "programs.json read error")
		return
	}

	var fileload map[string]any
	err = json.Unmarshal(data, &fileload)
	if err != nil {
		panic(err)
	}

	for _, v := range fileload["Programs"].([]any) {
		p := v.(map[string]any)
		var name = p["Name"].(string)
		var program = p["Program"].(string)
		var key = p["APIKey"].(string)
		var dir = p["Dir"].(string)
		var args = make([]string, len(p["Arguments"].([]any)))
		for count, v := range p["Arguments"].([]any) {
			args[count] = v.(string)
		}
		if program != "" && key != "" {
			api.Programs = append(api.Programs, api.Program{Name: name, Program: program, Dir: dir, APIKey: key, Arguments: args})
		} else {
			util.Log(util.MAIN, "Invalid Program:", program, key, args)
		}
	}

	util.Log(util.MAIN, "loaded Programs:")
	for _, v := range api.Programs {
		fmt.Printf("%s -> %s %s Key:%s \n", v.Name, v.Program, v.Arguments, v.APIKey)
	}
	fmt.Println()

	router := mux.NewRouter().StrictSlash(true)
	api.CreateAPI(router)
	util.Log(util.MAIN, "Started API")

	for _, v := range api.Programs {
		err = api.Register(&v)
		if err != nil {
			util.Err(util.MAIN, err, true, "Registering Error")
		}
	}

	// Blocking
	err = http.ListenAndServe(":"+fmt.Sprintf("%d", util.GetConfig().Port), router)
	util.Err(util.MAIN, err, true, "Listening Error")
}
