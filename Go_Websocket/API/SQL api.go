package api

import (
	_ "context"
	"database/sql"
	_ "fmt"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

func ProcessLogRequest(recive *map[string]interface{}) (bool, error) {
	return true, nil
}

func ProcessActivityRequest(recive *map[string]interface{}) (bool, error) {
	return true, nil
}
