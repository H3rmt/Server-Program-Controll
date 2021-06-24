package main

import (
	_ "fmt"
	_ "log"

	_ "context"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func checkadmin(js *map[string]interface{}) bool {
	code, code_exists := (*js)["code"]
	valid := code_exists && code == "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" //test
	return valid
}

// Permissionerror Exported
type Permissionerror struct{}

func (m *Permissionerror) Error() string {
	return "no admin permissions"
}

// Start Exported
func Start(recive *map[string]interface{}) (bool, error) {
	if !checkadmin(recive) {
		return false, &Permissionerror{}
	}
	return false, nil
}

// Stop Exported
func Stop(recive *map[string]interface{}) (bool, error) {
	if !checkadmin(recive) {
		return false, &Permissionerror{}
	}
	return false, nil
}

// Customaction Exported
func Customaction(recive *map[string]interface{}) (map[string]interface{}, error) {
	if !checkadmin(recive) {
		return nil, &Permissionerror{}
	}
	return map[string]interface{}{"success": false, "return": "raawr"}, nil
}
