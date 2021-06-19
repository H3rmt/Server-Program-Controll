package ExternalCommunication

import (
	_ "fmt"
	_ "log"

	_ "context"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func checkadmin(js *map[string]interface{}) bool {
	code, code_exists := (*js)["code"]
	valid := code_exists && code == "test" // TODO: implement complex check of hashed code
	return valid
}

// Exported
func Start(recive *map[string]interface{}) ([]interface{}, error) {
	checkadmin(recive)
	return nil, nil
}

// Exported
func Stop(recive *map[string]interface{}) ([]interface{}, error) {
	checkadmin(recive)
	return nil, nil
}

// Exported
func Customaction(recive *map[string]interface{}) ([]interface{}, error) {
	checkadmin(recive)
	return nil, nil
}
