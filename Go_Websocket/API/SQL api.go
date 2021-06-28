package api

import (
	ws "Go_Websocket/WS"
	"database/sql"
	"log"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

/*
finds the corresponding program with ID from the database
*/
func getProgramm_IDfromAPIKey(APIKey string) (string, error) {
	sql := "SELECT ID from programs WHERE APIKey=?;"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	// Execute query
	res, err := stmt.Query(APIKey)
	if err != nil {
		return "", err
	}

	if res.Next() {
		ID := ""
		res.Scan(&ID)
		return ID, nil
	} else {
		return "", &APIKeyerror{}
	}
}

/*
Error thrown/returned when no admin priv are present
*/
type APIKeyerror struct{}

func (m *APIKeyerror) Error() string {
	return "Invalid APIKey"
}

/*
adds log do the database from logrequest
*/
func ProcessLogRequest(APIKey string, logrequest *LogRequest) error {
	Programm_ID, err := getProgramm_IDfromAPIKey(APIKey)
	if err != nil {
		return err
	}

	sql := "INSERT INTO logs (Programm_ID,Date,Number,Message,Type) VALUES (?,?,?,?,?);"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Query(Programm_ID, (*logrequest).Date, (*logrequest).Number, (*logrequest).Message, (*logrequest).Type)
	if err != nil {
		return err
	} else {
		log.Println("Log added to database")
	}

	return nil
}

/*
Struct to represent a Request asking to add a log in the Log table in DB
*/
type LogRequest struct {
	Date    string
	Number  float64
	Message string
	Type    ws.Logtype
}

/*
adds activity do the database from activityrequest
*/
func ProcessActivityRequest(APIKey string, activityrequest *ActivityRequest) error {
	Programm_ID, err := getProgramm_IDfromAPIKey(APIKey)
	if err != nil {
		return err
	}

	sql := "INSERT INTO activity (Programm_ID,Date,Type) VALUES (?,?,?);"
	stmt, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Query(Programm_ID, (*activityrequest).Date, (*activityrequest).Type)
	if err != nil {
		return err
	} else {
		log.Println("Activity added to database")
	}

	return nil
}

/*
Struct to represent a Request asking to add a activity in the Acitivity table in DB
*/
type ActivityRequest struct {
	Date string
	Type ws.Activitytype
}
