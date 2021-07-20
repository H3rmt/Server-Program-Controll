package main

import (
	"fmt"
	"log"
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
	return &Program{}, &InvalidAPIkeyerror{}
}

/*
Process request to execute command in Program
*/
func ProcessCommandRequest(program *Program, request *ActualCommandRequest) (err error) {
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

/*
Struct to represent a Request send to actual program asking to execute command
*/
type ActualCommandRequest struct {
	Message string
	APIkey  string
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
