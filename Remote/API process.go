package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
finds the corresponding program from list of programs
*/
func getProgramm_IDfromAPIKey(APIKey string) (*Program, error) {
	for i := 0; i < len(programs); i++ {
		if programs[i].APIKey == APIKey {
			log.Println("Program found:", APIKey, programs[i].Program, programs[i].Arguments)
			return &programs[i], nil
		}
	}
	return &Program{}, &InvalidAPIKeyerror{}
}

/*
Process request to execute command in Program
*/
func ProcessCommandRequest(program *Program, request *CommandRequest) (err error) {
	switch request.Message {
	case "Start":
		err = program.Start()
	case "Stop":
		err = program.Stop()
	default:
		err = fmt.Errorf("unsuported Comand")
	}
	return
}

func SendLog(message string, program *Program, logtype Logtype) error {
	date := time.Now().Format("2006-01-02 15:04:05")
	logrequest := map[string]interface{}{"APIKey": program.APIKey, "Log": LogRequest{Date: date, Number: program.logcounter, Message: message, Type: logtype}}
	jsonReq, err := json.Marshal(logrequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(remoteIP+":18769/api", "application/json;", bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println("recived answer:", string(bodyBytes))

	var answer Answer
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		return err
	}

	if answer.Success {
		return nil
	} else {
		return fmt.Errorf("not successfully added log")
	}

}

/*
Struct to represent a Answer from the program
*/
type Answer struct {
	Success bool
}

/*
Struct to represent a Request send to actual program asking to execute command
*/
type CommandRequest struct {
	Message string
	APIKey  string
}

/*
Struct to represent a Request asking to add a log in the Log table in DB
*/
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

/*
Struct to represent a Request asking to add a activity in the Acitivity table in DB
*/
type ActivityRequest struct {
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
