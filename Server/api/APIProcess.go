package api

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"Server/util"
)

var ProgramConnections = make(map[string]string) // Program_ID = [IP:port]

var DB *sqlx.DB

func SetDB(db *sqlx.DB) {
	DB = db
}

/*
Process request to add Program to list of connections
*/
func ProcessRegisterRequest(Progamm_ID string, addr string, port uint16) error {
	if _, exist := ProgramConnections[Progamm_ID]; exist {
		util.Log("PRGR API", addr, ":", port, " relinked with id:", Progamm_ID)
	} else {
		util.Log("PRGR API", addr, ":", port, " linked with id:", Progamm_ID)
	}
	ProgramConnections[Progamm_ID] = fmt.Sprintf("%s:%d", addr, port)
	return nil
}

/*
Error returned when APIKey was invalid
*/
type SQLerror struct{}

func (m *SQLerror) Error() string {
	return "SQLerror"
}

/*
finds the corresponding program with ID from the database
*/
func getProgram_IDfromAPIKey(APIKey string) (string, error) {
	sql := "SELECT ID from programs WHERE APIKey=?"

	setting := struct {
		ID string `db:"ID"`
	}{}
	err := DB.Get(&setting, sql, APIKey)
	if err != nil {
		util.Log("PRGR API", err)
		return "", &SQLerror{}
	}

	if setting.ID != "" {
		return setting.ID, nil
	} else {
		return "", &InvalidAPIKeyerror{}
	}
}

/*
Error returned when APIKey was invalid
*/
type InvalidAPIKeyerror struct{}

func (m *InvalidAPIKeyerror) Error() string {
	return "Invalid API key"
}

/*
adds log do the database from logrequest
*/
func ProcessLogRequest(Program_ID string, logrequest *LogRequest) error {
	sq := "INSERT INTO logs (program_ID,Date,Number,Message,Type) VALUES (:programId, :Date, :Number, :Message, :Type)"
	_, err := DB.NamedExec(sq, &struct {
		ProgramId string `db:"programId"`
		Date      string `db:"Date"`
		Message   string `db:"Message"`
		Type      string `db:"Type"`
	}{Program_ID, logrequest.Date, logrequest.Message, string(logrequest.Type)})
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	} else {
		util.Log("SQL API", "Log added to database")
	}
	return nil
}

/*
Struct to represent a Request asking to add a log in the Log table in DB
*/
type LogRequest struct {
	Date    string
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
adds activity do the database from activityrequest
*/
func ProcessActivityRequest(Program_ID string, activityrequest *ActivityRequest) error {
	sq := "INSERT INTO activity (program_ID,Date,Type) VALUES (:programId, :Date, :Type)"
	_, err := DB.NamedExec(sq, &struct {
		ProgramId string `db:"programId"`
		Date      string `db:"Date"`
		Type      string `db:"Type"`
	}{Program_ID, activityrequest.Date, string(activityrequest.Type)})
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	} else {
		util.Log("SQL API", "Activity added to database")
	}
	return nil
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
Process request telling that the program stopped
*/
func ProcessStateChangeRequest(Program_ID string, statechangerequest *StateChangeRequest) error {
	sq := "UPDATE programs SET Active = ?, StatechangeTime = ? WHERE ID = ?;"

	_, err := DB.Exec(sq, statechangerequest.Start, statechangerequest.Date, Program_ID)
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	}

	sq = "INSERT INTO logs (program_ID,Date,Number,Message,Type) VALUES (:program_ID, :Date, :Number, :Message, :Type)"

	message := ""
	if statechangerequest.Start {
		message = "START"
	} else {
		message = "STOP"
	}

	_, err = DB.NamedExec(sq, &struct {
		ProgramId string `db:"program_ID"`
		Date      string `db:"Date"`
		Message   string `db:"Message"`
		Type      string `db:"Type"`
	}{Program_ID, statechangerequest.Date, message, "Important"})
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	} else {
		util.Log("SQL API", "Log added to database")
	}

	return nil
}

/*
Struct to represent a Request telling that the program stopped or started
*/
type StateChangeRequest struct {
	Date  string
	Start bool
}
