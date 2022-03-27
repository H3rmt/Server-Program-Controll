package websocket

import (
	"github.com/jmoiron/sqlx"

	"Server/util"
)

var DB *sqlx.DB

func SetDB(db *sqlx.DB) {
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
check if send shacode exists or equals stored sha code
*/
func Checkadmin(js *map[string]any) error {
	code, code_exists := (*js)["admin"]
	if code_exists {
		sql := "SELECT Value FROM settings WHERE Name='adminCookie'"

		setting := struct {
			Value string `db:"Value"`
		}{""}
		err := DB.Get(&setting, sql)
		if err != nil {
			util.Log("SQL WS", err)
			return &SQLerror{}
		}

		if setting.Value == code {
			return nil
		}
		return &Permissionerror{}

	}
	return &Permissionerror{}
}

/*
finds the corresponding program with ID from the database
*/
func getAPIKeyfromProgram_ID(Program_ID string) (string, error) {
	sql := "SELECT APIKey from programs WHERE ID=?"

	setting := struct {
		APIKey string `db:"APIKey"`
	}{}
	err := DB.Get(&setting, sql, Program_ID)
	if err != nil {
		util.Log("SQL WS", err)
		return "", &SQLerror{}
	}

	if setting.APIKey != "" {
		return setting.APIKey, nil
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
returns logs as array of Log-structs from DB (logs)
*/
func Getlogs(Program_id string) ([]Log, error) {
	sql := "SELECT Date,Number,Message,Type FROM logs WHERE program_ID=?"

	// []Log{} so it sends empty array instead of null
	var logs = []Log{}
	err := DB.Select(&logs, sql, Program_id)
	if err != nil {
		util.Log("SQL WS", err)
		return nil, &SQLerror{}
	}
	return logs, nil
}

/*
Struct to represent a row in the Log table
*/
type Log struct {
	Date    string  `db:"Date"`
	Number  float64 `db:"Number"`
	Message string  `db:"Message"`
	Type    Logtype `db:"Type"`
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
func Getactivity(Program_id string) ([]Activity, error) {
	sql := "SELECT Date,Type FROM activity WHERE program_ID=?"

	// []Activity{} so it sends empty array instead of null
	var activities = []Activity{}
	err := DB.Select(&activities, sql, Program_id)
	if err != nil {
		util.Log("SQL WS", err)
		return nil, &SQLerror{}
	}
	return activities, nil
}

/*
Struct to represent a row in the Activity table
*/
type Activity struct {
	Date string       `db:"Date"`
	Type Activitytype `db:"Type"`
}

type Activitytype string

const (
	Send              Activitytype = "Send"
	Recive            Activitytype = "Recive"
	Process           Activitytype = "Process"
	Backgroundprocess Activitytype = "Backgroundprocess"
)
