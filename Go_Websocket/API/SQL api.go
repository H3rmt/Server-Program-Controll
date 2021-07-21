package api

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

/*
finds the corresponding program with ID from the database
*/
func getProgramm_IDfromAPIKey(APIKey string) (string, error) {
	sql := "SELECT ID from programs WHERE APIKey=?;"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	// Execute query
	res, err := stmt.Query(APIKey)
	if err != nil {
		return "", err
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
func ProcessLogRequest(Programm_ID string, logrequest *LogRequest) error {
	sql := "INSERT INTO logs (Programm_ID,Date,Number,Message,Type) VALUES (?,?,?,?,?);"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Query(Programm_ID, (*logrequest).Date, (*logrequest).Number, (*logrequest).Message, (*logrequest).Type)
	if err != nil {
		return err
	} else {
		log.Println("SQL API|", "Log added to database")
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
func ProcessActivityRequest(Programm_ID string, activityrequest *ActivityRequest) error {
	sql := "INSERT INTO activity (Programm_ID,Date,Type) VALUES (?,?,?);"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Query(Programm_ID, (*activityrequest).Date, (*activityrequest).Type)
	if err != nil {
		return err
	} else {
		log.Println("SQL API|", "Activity added to database")
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
