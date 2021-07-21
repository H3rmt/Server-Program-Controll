package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	api "Go_Websocket/API"
)

var Port = "18770"

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

func ProcessCommand(Program_id string, command string) error {
	IP, exists := api.Programconnections[Program_id]
	if !exists {
		log.Println("PRGR API|", "IP not registered", Program_id)
		return &Programerror{"IP not registered"}
	}

	APIkey, err := getAPIKeyfromProgramm_ID(Program_id)
	if err != nil {
		log.Println("PRGR API|", "Programm_ID err", err, Program_id)
		return err
	}

	request := CommandRequest{command, APIkey}

	byterequest, err := json.Marshal(request)
	if err != nil {
		log.Println("PRGR API|", "Requestbuild invalied", err, request)
		return err
	}

	resp, err := http.Post(fmt.Sprintf("http://%s:%s/api", IP, Port), "application/json;", bytes.NewBuffer(byterequest))
	if err != nil {
		log.Println("PRGR API|", "Program did not respond", err, IP+":"+Port)
		return &Programerror{"Program did not respond"}
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("PRGR API|", "Program response invalid", err, resp.Body)
		return &Programerror{"Program respond invalid"}
	}

	log.Println("PRGR API|", "recived answer:", string(bodyBytes))

	var answer CommandAnswer
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		log.Println("PRGR API|", "Program Json response invalid", err)
		return &Programerror{"Program Json response invalid"}
	}

	if answer.Success {
		return nil
	} else {
		return &Programerror{"Command not successfully executed"}
	}
}

/*
Error thrown/returned when no admin priviges are present
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
	APIkey  string
}

/*
Struct to represent a Answer from the program having executed a command
*/
type CommandAnswer struct {
	Success bool
	APIkey  string
}
