package api

import (
	"database/sql"

	"Server/util"
)

var Programconnections = make(map[string]string) // Program_ID = [IP]

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

/*
Process request to add Program to list of connections
*/
func ProcessRegisterRequest(Progamm_ID string, addr string) error {
	if _, exist := Programconnections[Progamm_ID]; exist {
		util.Log("PRGR API", addr, " relinked with id:", Progamm_ID)
	} else {
		util.Log("PRGR API", addr, " linked with id:", Progamm_ID)
	}
	Programconnections[Progamm_ID] = addr
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
	sql := "SELECT ID from programs WHERE APIKey=?;"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		util.Log("PRGR API", err)
		return "", &SQLerror{}
	}
	defer stmt.Close()

	// Execute query
	res, err := stmt.Query(APIKey)
	if err != nil {
		util.Log("PRGR API", err)
		return "", &SQLerror{}
	}

	if res.Next() {
		ID := ""
		res.Scan(&ID)
		return ID, nil
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
	sql := "INSERT INTO logs (program_ID,Date,Number,Message,Type) VALUES (?,?,?,?,?);"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Query(Program_ID, logrequest.Date, logrequest.Number, logrequest.Message, logrequest.Type)
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
	Number  float64
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
	sql := "INSERT INTO activity (program_ID,Date,Type) VALUES (?,?,?);"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Query(Program_ID, activityrequest.Date, activityrequest.Type)
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
	sql := "UPDATE programs SET Active = ?, StartStoptime = ? WHERE ID = ?;"

	stmt, err := DB.Prepare(sql)
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	}

	// Execute query
	_, err = stmt.Query(statechangerequest.Start, statechangerequest.Date, Program_ID)
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	}

	stmt.Close()

	sql = "INSERT INTO logs (program_ID,Date,Number,Message,Type) VALUES (?,?,?,?,?);"

	stmt, err = DB.Prepare(sql)
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	}

	message := ""
	if statechangerequest.Start {
		message = "START"
	} else {
		message = "STOP"
	}

	logtype := ""
	if statechangerequest.Start {
		logtype = "Important"
	} else {
		logtype = "Error"
	}

	// Execute query
	_, err = stmt.Query(Program_ID, statechangerequest.Date, statechangerequest.Number, message, logtype)
	if err != nil {
		util.Log("PRGR API", err)
		return &SQLerror{}
	} else {
		util.Log("SQL API", "State Change added to database")
	}

	stmt.Close()

	return nil
}

/*
Struct to represent a Request telling that the program stopped or started
*/
type StateChangeRequest struct {
	Date   string
	Number int
	Start  bool
}
