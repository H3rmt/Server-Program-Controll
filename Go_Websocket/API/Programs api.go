package api

import (
	"log"
)

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
Struct to represent a StateChange of a Program
*/
type StateChangeRequest struct {
	State State
}

type State bool

const (
	Running State = true
	Stopped State = false
)
