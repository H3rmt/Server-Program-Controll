package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var Port = "18769"

/*

 */
func main() {
	log.Println("Performing Http Post...")
	req := map[string]interface{}{"APIkey": "5rrtg3u564uiqr43fadf", "LogRequest": LogRequest{Date: "2021-06-17 22:26:43", Number: 12, Message: "hitext", Type: Important}}
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
}

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

type Activity struct {
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
