package ws

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

/*
finds the corresponding program with ID from the database
*/
func getAPIKeyfromProgramm_ID(Programm_ID string) (string, error) {
	sql := "SELECT APIKey from programs WHERE ID=?;"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	// Execute query
	res, err := stmt.Query(Programm_ID)
	if err != nil {
		return "", err
	}

	if res.Next() {
		ID := ""
		res.Scan(&ID)
		return ID, nil
	} else {
		return "", &InvalidAPIkeyerror{}
	}
}

/*
Error returned when APIkey was invalid
*/
type InvalidAPIkeyerror struct{}

func (m *InvalidAPIkeyerror) Error() string {
	return "Invalid API key"
}

/*
returns how many rows will be returned from table
*/
func getRowcount(table string) int {
	sql := fmt.Sprintf("SELECT COUNT(*) FROM %s;", table)

	// Execute query
	query, err := DB.Query(sql)
	if err != nil {
		return 0
	}
	defer query.Close()

	var count int
	query.Scan(&count)

	return count
}

/*
returns logs as array of Log-structs from DB (logs)
*/
func Getlogs(Program_id string) ([]interface{}, error) {
	count := getRowcount("logs")

	var entries = make([]interface{}, count)

	sql := "SELECT programm_ID,Date,Number,Message,Type FROM logs WHERE programm_ID=?"

	stmt, err := DB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(Program_id)
	if err != nil {
		return nil, err
	}

	// iterate through query
	for rows.Next() {
		var Programm_ID string
		var Date, Message string
		var Number float64
		var Type Logtype

		// Get values from row
		err := rows.Scan(&Programm_ID, &Date, &Number, &Message, &Type)
		if err != nil {
			return nil, err
		}
		entries = append(entries, Log{Programm_ID: Programm_ID, Date: Date, Number: Number, Message: Message, Type: Type})
	}
	return entries, nil
}

/*
Struct to represent a row in the Log table
*/
type Log struct {
	Programm_ID string
	Date        string
	Number      float64
	Message     string
	Type        Logtype
}

type Logtype string

const (
	Low       Logtype = "Low"
	Normal    Logtype = "Normal"
	Important Logtype = "Important"
	Error     Logtype = "Error"
)

/*
returns activity as array of Activity-structs from DB (activity)
*/
func Getactivity(Program_id string) ([]interface{}, error) {
	count := getRowcount("activity")

	var entries = make([]interface{}, count)

	sql := "SELECT programm_ID,Date,Type FROM activity WHERE programm_ID=?"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(Program_id)
	if err != nil {
		return nil, err
	}

	// iterate through query
	for rows.Next() {
		var Programm_ID string
		var Date string
		var Type Activitytype

		// get values from row
		err := rows.Scan(&Programm_ID, &Date, &Type)
		if err != nil {
			return nil, err
		}
		entries = append(entries, Activity{Programm_ID: Programm_ID, Date: Date, Type: Type})
	}
	return entries, nil
}

/*
Struct to represent a row in the Activity table
*/
type Activity struct {
	Programm_ID string
	Date        string
	Type        Activitytype
}

type Activitytype string

const (
	Send              Activitytype = "Send"
	Recive            Activitytype = "Recive"
	Process           Activitytype = "Process"
	Backgroundprocess Activitytype = "Backgroundprocess"
)
