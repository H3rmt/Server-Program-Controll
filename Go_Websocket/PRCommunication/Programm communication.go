package PRCommunication

// Exported
func checkadmin(js *map[string]interface{}) bool {
	code, code_exists := (*js)["code"]
	valid := code_exists && code == "test"
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
