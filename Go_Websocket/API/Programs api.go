package api

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "context"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Port = "18770"

var Programconnections = make(map[string]string) // Programm_ID = [IP]

/*
Process request to add Program to list of connections
*/
func ProcessRegisterRequest(ProgammID string, addr string) error {
	if _, exist := Programconnections[ProgammID]; exist {
		log.Println("PRGR API|", addr, " relinked with id:", ProgammID)
	} else {
		log.Println("PRGR API|", addr, " linked with id:", ProgammID)
	}
	Programconnections[ProgammID] = addr
	return nil
}

/*
Process request to execute command in Program
*/
func ProcessCommandRequest(ProgammID string, commandrequest *CommandRequest, APIkey string) error {
	IP, exists := Programconnections[ProgammID]
	if !exists {
		log.Println("PRGR API|", "IP not registered")
		return &Programnotreachableerror{"IP not registered"}
	}

	request := ActualCommandRequest{commandrequest.Message, APIkey}

	byterequest, err := json.Marshal(request)
	if err != nil {
		log.Println("PRGR API|", "Requestbuild invalied")
		return err
	}

	resp, err := http.Post(IP+":"+Port, "application/json;", bytes.NewBuffer(byterequest))
	if err != nil {
		log.Println("PRGR API|", "Program did not respond")
		return &Programnotreachableerror{"Program did not respond"}
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("PRGR API|", "Program response invalid")
		return &Programnotreachableerror{"Program respond invalid"}
	}

	log.Println("PRGR API|", "recived answer:", string(bodyBytes))

	var answer ActualCommandAnswer
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		log.Println("PRGR API|", "Program Json response invalid")
		return &Programnotreachableerror{"Program Json response invalid"}
	}

	if answer.Success {
		return nil
	} else {
		return &Programnotreachableerror{"Command not successfully executed"}
	}

}

/*
Struct to represent a Request asking to execute command in Program
*/
type CommandRequest struct {
	Message string
}

/*
Struct to represent a Request send to actual program asking to execute command
*/
type ActualCommandRequest struct {
	Message string
	APIkey  string
}

/*
Struct to represent a Answer from the actual program having executed a command
*/
type ActualCommandAnswer struct {
	Success bool
	APIkey  string
}

/*
Error thrown/returned when no admin priviges are present
*/
type Programnotreachableerror struct {
	message string
}

func (m *Programnotreachableerror) Error() string {
	return m.message
}
