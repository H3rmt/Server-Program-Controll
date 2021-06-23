package SQL

import (
	"fmt"
	"log"

	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var user = "Go"
var password = "e73EG6dP2f8F2dAx"
var database = "programs"

// Init Exported
func Init() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, database))
	if err != nil {
		log.Fatal("Error creating connection: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Connected!\n")
}

func getRowcount(database string) int {
	ctx := context.Background()

	sql := fmt.Sprintf("SELECT COUNT(*) FROM %s;", database)

	// Execute query
	query, err := db.QueryContext(ctx, sql)
	if err != nil {
		return 0
	}
	defer query.Close()

	var count int
	query.Scan(&count)

	return count
}

// Getlogs Exported
func Getlogs(recive *map[string]interface{}) ([]interface{}, error) {
	ctx := context.Background()

	count := getRowcount("logs")

	var entries = make([]interface{}, count)

	sql := "SELECT programm_ID,Date,Number,Message,Type FROM logs;"

	// Execute query
	query, err := db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// Iterate through the result set.
	for query.Next() {
		var Programm_ID string
		var Date, Message string
		var Number int
		var Type Logtype

		// Get values from row.
		err := query.Scan(&Programm_ID, &Date, &Number, &Message, &Type)
		if err != nil {
			return nil, err
		}
		entries = append(entries, Log{Programm_ID: Programm_ID, Date: Date, Number: Number, Message: Message, Type: Type})
	}
	return entries, nil
}

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

// Getactivity Exported
func Getactivity(recive *map[string]interface{}) ([]interface{}, error) {
	ctx := context.Background()

	count := getRowcount("activity")

	var entries = make([]interface{}, count)

	sql := "SELECT programm_ID,Date,Type FROM activity;"

	// Execute query
	query, err := db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// Iterate through the result set.
	for query.Next() {
		var Programm_ID string
		var Date string
		var Type Logtype

		// Get values from row.
		err := query.Scan(&Programm_ID, &Date, &Type)
		if err != nil {
			return nil, err
		}
		entries = append(entries, Log{Programm_ID: Programm_ID, Date: Date, Type: Type})
	}
	return entries, nil
}

type Activity struct {
	Programm_ID string
	Date        string
	Number      int
	Message     string
	Type        Activitytype
}

type Activitytype string

const (
	Send              Activitytype = "Send"
	Recive            Activitytype = "Recive"
	Process           Activitytype = "Process"
	Backgroundprocess Activitytype = "Backgroundprocess"
)
