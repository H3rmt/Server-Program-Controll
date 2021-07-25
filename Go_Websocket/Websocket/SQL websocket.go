package ws

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
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
func getAPIKeyfromProgram_ID(Program_ID string) (string, error) {
	sql := "SELECT APIKey from programs WHERE ID=?;"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Println(err)
		return "", &SQLerror{}
	}
	defer stmt.Close()

	// Execute query
	res, err := stmt.Query(Program_ID)
	if err != nil {
		log.Println(err)
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

	sql := "SELECT Date,Number,Message,Type FROM logs WHERE program_ID=?"

	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil, &SQLerror{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(Program_id)
	if err != nil {
		log.Println(err)
		return nil, &SQLerror{}
	}

	// iterate through query
	for rows.Next() {
		var Date, Message string
		var Number float64
		var Type Logtype

		// Get values from row
		err := rows.Scan(&Date, &Number, &Message, &Type)
		if err != nil {
			log.Println(err)
			return nil, &SQLerror{}
		}
		entries = append(entries, Log{Date: Date, Number: Number, Message: Message, Type: Type})
	}
	return entries, nil
}

/*
Struct to represent a row in the Log table
*/
type Log struct {
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
returns activity as array of Activity-structs from DB (activity)
*/
func Getactivity(Program_id string) ([]interface{}, error) {
	count := getRowcount("activity")

	var entries = make([]interface{}, count)

	sql := "SELECT Date,Type FROM activity WHERE program_ID=?"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil, &SQLerror{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(Program_id)
	if err != nil {
		log.Println(err)
		return nil, &SQLerror{}
	}

	// iterate through query
	for rows.Next() {
		var Date string
		var Type Activitytype

		// get values from row
		err := rows.Scan(&Date, &Type)
		if err != nil {
			log.Println(err)
			return nil, &SQLerror{}
		}
		entries = append(entries, Activity{Date: Date, Type: Type})
	}
	return entries, nil
}

/*
Struct to represent a row in the Activity table
*/
type Activity struct {
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
