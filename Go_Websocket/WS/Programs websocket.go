package ws

import (
	_ "fmt"
	_ "log"

	_ "context"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/*
check if send shacode exists or equals stored sha code
*/
func checkadmin(js *map[string]interface{}) bool {
	code, code_exists := (*js)["code"]
	valid := code_exists && code == "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" //test
	return valid
}

/*
Error thrown/returned when no admin priv are present
*/
type Permissionerror struct{}

func (m *Permissionerror) Error() string {
	return "no admin permissions"
}

/*
returns true if program was successfully started
*/
func Start(recive *map[string]interface{}) (bool, error) {
	if !checkadmin(recive) {
		return false, &Permissionerror{}
	}
	return false, nil
}

/*
returns true if program was successfully stoped
*/
func Stop(recive *map[string]interface{}) (bool, error) {
	if !checkadmin(recive) {
		return false, &Permissionerror{}
	}
	return false, nil
}

/*
returns map with values from program if custom comand was executed successfully
else map with false and error codes
*/
func Customaction(recive *map[string]interface{}) (map[string]interface{}, error) {
	if !checkadmin(recive) {
		return nil, &Permissionerror{}
	}
	return map[string]interface{}{"success": false, "return": "raawr"}, nil
}
