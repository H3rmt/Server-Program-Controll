package api

import (
	_ "fmt"
	"log"

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
func ProcessCommandRequest(ProgammID string, customcommandrequest *CommandRequest) error {
	return nil
}

/*
Struct to represent a Request asking to execute custom command in Program
*/
type CommandRequest struct {
	Message string
}
