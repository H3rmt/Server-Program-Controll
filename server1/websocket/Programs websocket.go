package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"Server/api"
	"Server/util"
)

/*
returns true if program was successfully started
*/
func Start(Program_id string) (map[string]string, error) {
	err := ProcessCommand(Program_id, "Start")
	if err != nil {
		return map[string]string{"success": "false"}, err
	} else {
		return map[string]string{"success": "true"}, nil
	}
}

/*
returns true if program was successfully stoped
*/
func Stop(Program_id string) (map[string]string, error) {
	err := ProcessCommand(Program_id, "Stop")
	if err != nil {
		return map[string]string{"success": "false"}, err
	} else {
		return map[string]string{"success": "true"}, nil
	}
}

/*
Sends a command to the program using the registered IP address

throws Programerrors if error happen
*/
func ProcessCommand(Program_id string, command string) error {
	IPPort, exists := api.Programconnections[Program_id]
	if !exists {
		util.Log("PRGR WS", "IP not registered", Program_id)
		return &Programerror{"IP not registered"}
	}

	APIKey, err := getAPIKeyfromProgram_ID(Program_id)
	if err != nil {
		util.Log("PRGR WS", "Program_ID err", err, Program_id)
		return err
	}

	request := CommandRequest{command, APIKey}

	byterequest, err := json.Marshal(request)
	if err != nil {
		util.Log("PRGR WS", "Requestbuild invalied", err, request)
		return err
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/api", IPPort), "application/json;", bytes.NewBuffer(byterequest))
	if err != nil {
		util.Log("PRGR WS", "Program did not respond", err, IPPort)
		return &Programerror{"Program did not respond"}
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.Log("PRGR WS", "Program response invalid", err, resp.Body)
		return &Programerror{"Program respond invalid"}
	}

	util.Log("PRGR WS", "recived answer:", string(bodyBytes))

	var answer Answer
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		util.Log("PRGR WS", "Program Json response invalid", err)
		return &Programerror{"Program Json response invalid"}
	}

	if answer.Success {
		return nil
	} else {
		util.Log("PRGR WS", "Command not successfully executed")
		return &Programerror{"Command not successfully executed"}
	}
}

/*
Error thrown/returned when error while communicating with the program occured
*/
type Programerror struct {
	message string
}

func (m *Programerror) Error() string {
	return m.message
}

/*
Struct to represent a Request send to program asking to execute command
*/
type CommandRequest struct {
	Message string
	APIKey  string
}

/*
Struct to represent a Answer from the program
*/
type Answer struct {
	Success bool
}
