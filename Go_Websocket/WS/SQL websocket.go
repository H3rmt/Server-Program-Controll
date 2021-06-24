package ws

import (
	"context"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

/*
returns how many rows will be returned from table
*/
func getRowcount(table string) int {
	ctx := context.Background()

	sql := fmt.Sprintf("SELECT COUNT(*) FROM %s;", table)

	// Execute query
	query, err := DB.QueryContext(ctx, sql)
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
func Getlogs(recive *map[string]interface{}) ([]interface{}, error) {
	ctx := context.Background()

	count := getRowcount("logs")

	var entries = make([]interface{}, count)

	sql := "SELECT programm_ID,Date,Number,Message,Type FROM logs;"

	// Execute query
	query, err := DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// iterate through query
	for query.Next() {
		var Programm_ID string
		var Date, Message string
		var Number int
		var Type Logtype

		// Get values from row
		err := query.Scan(&Programm_ID, &Date, &Number, &Message, &Type)
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
	Number      int
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
func Getactivity(recive *map[string]interface{}) ([]interface{}, error) {
	ctx := context.Background()

	count := getRowcount("activity")

	var entries = make([]interface{}, count)

	sql := "SELECT programm_ID,Date,Type FROM activity;"

	// execute query
	query, err := DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// iterate through query
	for query.Next() {
		var Programm_ID string
		var Date string
		var Type Activitytype

		// get values from row
		err := query.Scan(&Programm_ID, &Date, &Type)
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
