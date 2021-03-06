package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Port            uint16
	LogPrefix       bool
	Prefixstretch   int8
	Locationstretch int8
	User            string
	Password        string
	Database        string
	Database2       string
	DbPort          uint16
}

var conf config

func GetConfig() *config {
	return &conf
}

func LoadConfig() error {
	defaultConfig()

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		Err(CONFIG, err, true, "Error reading", "Config/config.json", "file")
		return err
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		Err(CONFIG, err, true, "Error unmarshalling configs")
		return err
	}
	Log(CONFIG, "Loaded config:", fmt.Sprintf("%+v", conf))
	return nil
}

func defaultConfig() {
	conf.Port = 0
	conf.LogPrefix = true
	conf.Prefixstretch = 0
	conf.Locationstretch = 0
	conf.User = ""
	conf.Password = ""
	conf.Database = ""
	conf.Database2 = ""
	conf.DbPort = 0
}
