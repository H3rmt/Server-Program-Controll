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

var Programconnections map[string]string // Programm_ID = [IP]

/*
Process request to add Program to list of connections
*/
func ProcessRegisterRequest(ProgammID string, registerrequest *RegisterRequest) error {
	if _, exist := Programconnections[ProgammID]; exist {
		log.Println(registerrequest.IP, " linked with ", ProgammID)
	} else {
		log.Println(registerrequest.IP, " relinked with ", ProgammID)
	}
	Programconnections[ProgammID] = registerrequest.IP
	return nil
}

/*
Struct to represent a Request registering itself
*/
type RegisterRequest struct {
	IP string
}

/*
Process request to execute custom command in Program
*/
func ProcessCommandRequest(ProgammID string, customcommandrequest *CommandRequest, APIkey string) error {
	IP, exists := Programconnections[ProgammID]
	if !exists {
		log.Println("IP not registered")
		return &Programnotreachableerror{"IP not registered"}
	}

	request := ActualCommandRequest{customcommandrequest.Message, APIkey}

	byterequest, err := json.Marshal(request)
	if err != nil {
		log.Println("Requestbuild invalied")
		return err
	}

	resp, err := http.Post(IP+":18769", "application/json;", bytes.NewBuffer(byterequest))
	if err != nil {
		log.Println("Program did not respond")
		return &Programnotreachableerror{"Program did not respond"}
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Program response invalid")
		return &Programnotreachableerror{"Program respond invalid"}
	}

	log.Println("recived answer:", string(bodyBytes))

	var answer ActualCommandAnswer
	err = json.Unmarshal(bodyBytes, &answer)
	if err != nil {
		log.Println("Program Json response invalid")
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
