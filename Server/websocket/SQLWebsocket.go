package websocket

import (
	"github.com/jmoiron/sqlx"

	"Server/util"
)

var PDB *sqlx.DB
var ADB *sqlx.DB

func SetPDB(db *sqlx.DB) {
	PDB = db
}

func SetADB(db *sqlx.DB) {
	ADB = db
}

/*
Error returned when APIKey was invalid
*/
type SQLerror struct{}

func (m *SQLerror) Error() string {
	return "SQLerror"
}

/*
CheckPermission checks if user is allowed to do action on program
*/
func CheckPermission(json *map[string]any) error {
	username, exists := (*json)["username"]
	hash, exists2 := (*json)["hash"]
	action, exists3 := (*json)["action"]
	id, exists4 := (*json)["id"]
	if exists && exists2 && exists3 && exists4 {
		sql := "SELECT admin, ID FROM users WHERE ID = (SELECT user_id FROM sessions WHERE user_id = (SELECT ID FROM users WHERE name = ?) AND hash = ?)"

		user := struct {
			Admin  bool `db:"admin"`
			UserId int  `db:"ID"`
		}{}
		err := ADB.Get(&user, sql, username.(string), hash.(string))
		if err != nil {
			util.Log("SQL WS", err, "Non-existing session")
			return &Permissionerror{}
		} // if no error happens session exists

		// admin has permission to do everything
		if user.Admin {
			return nil
		}

		// user is no admin => check program permissions
		sql = "SELECT permission FROM user_programs_permissions WHERE user_id = ? AND program_id = ?"
		permission := struct {
			Permission int `db:"permission"`
		}{}
		err = ADB.Get(&permission, sql, user.UserId, id.(float64))
		if err != nil {
			util.Log("SQL WS", err)
			return &Permissionerror{}
		}
		switch action {
		case "Logs":
			return nil // user is registered in user_programs_permissions => at least read access
		case "Activity":
			return nil // user is registered in user_programs_permissions => at least read access
		case "Start":
			if permission.Permission >= 1 {
				return nil
			}
		case "Stop":
			if permission.Permission >= 2 {
				return nil // user is registered in user_programs_permissions => at least read access
			}
		}
		return &Permissionerror{}
	}
	return &SQLerror{}
}

/*
finds the corresponding program with ID from the database
*/
func getAPIKeyfromProgram_ID(Program_ID string) (string, error) {
	sql := "SELECT APIKey from programs WHERE ID=?"

	setting := struct {
		APIKey string `db:"APIKey"`
	}{}
	err := PDB.Get(&setting, sql, Program_ID)
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
	sql := "SELECT Date,Message,Type FROM logs WHERE program_ID=?"

	// []Log{} so it sends empty array instead of null
	var logs = []Log{}
	err := PDB.Select(&logs, sql, Program_id)
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
	err := PDB.Select(&activities, sql, Program_id)
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
