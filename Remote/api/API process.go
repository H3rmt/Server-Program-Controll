package api

import (
	"Remote/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

/*
Struct to represent a Request send to actual program asking to execute command
*/
type CommandRequest struct {
	Message string
	APIKey  string
}

/*
Send Activity to server
*/
func SendActivity(program *Program, activitytype Activitytype) error {
	date := time.Now().Format("2006-01-02 15:04:05")
	activityrequest := map[string]interface{}{"APIKey": program.APIKey, "Activity": ActivityRequest{Date: date, Type: activitytype}}
	jsonReq, err := json.Marshal(activityrequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("http://%s:%d/api", util.GetConfig().RemoteIP, util.GetConfig().RemotePort), "application/json;", bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	util.Log("PRGR API", "recived answer:", string(bodyBytes))
	fmt.Println()

	var answer Answer
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		return err
	}

	if answer.Success {
		return nil
	} else {
		return fmt.Errorf("activity not added successfully")
	}

}

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

/*
Send Log to server
*/
func SendLog(message string, program *Program, logtype Logtype) error {
	date := time.Now().Format("2006-01-02 15:04:05")
	logrequest := map[string]interface{}{"APIKey": program.APIKey, "Log": LogRequest{Date: date, Number: program.logcounter, Message: message, Type: logtype}}
	jsonReq, err := json.Marshal(logrequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("http://%s:%d/api", util.GetConfig().RemoteIP, util.GetConfig().RemotePort), "application/json;", bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	util.Log("PRGR API", "recived answer:", string(bodyBytes))
	fmt.Println()

	var answer Answer
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		return err
	}

	if answer.Success {
		return nil
	} else {
		return fmt.Errorf("log not added successfully")
	}
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
Send State Change to server
*/
func SendStateChange(program *Program, Start bool) error {
	date := time.Now().Format("2006-01-02 15:04:05")
	statechangerequest := map[string]interface{}{"APIKey": program.APIKey, "StateChange": StateChangeRequest{Date: date, Number: program.logcounter, Start: Start}}
	jsonReq, err := json.Marshal(statechangerequest)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("http://%s:%d/api", util.GetConfig().RemoteIP, util.GetConfig().RemotePort), "application/json;", bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	util.Log("PRGR API", "recived answer:", string(bodyBytes))
	fmt.Println()

	var answer Answer
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		return err
	}

	if answer.Success {
		return nil
	} else {
		return fmt.Errorf("statechangerequest not transmitted successfully")
	}
}

/*
Struct to represent a Request telling that the program stopped or started
*/
type StateChangeRequest struct {
	Date   string
	Number int
	Start  bool
}

/*
Struct to represent a Answer from the program
*/
type Answer struct {
	Success bool
}
