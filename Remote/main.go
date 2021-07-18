package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

var Port = "18769"

/*

 */
func main() {
	cmd := exec.Command("python", "test.py")

	reader := Reader{}

	cmd.Stdout = &reader.outReader
	cmd.Stderr = &reader.errReader

	go reader.process()

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	cmd.Wait()
}

type Reader struct {
	outReader stdoutWriter
	errReader stderrWriter
}

func (r *Reader) processOutput(out string) {
	log.Println("out:", string(out))
}

func (r *Reader) processError(err string) {
	log.Println("err:", string(err))
}

func (r *Reader) process() {
	for {
		if r.outReader.buffer.Len() > 0 {
			read := make([]byte, r.outReader.buffer.Len())
			_, err := r.outReader.buffer.Read(read)
			if err != nil {
				log.Println("err in out reading", err)
			} else {
				r.processOutput(strings.TrimSpace(string(read)))
			}
		}
		if r.errReader.buffer.Len() > 0 {
			read := make([]byte, r.errReader.buffer.Len())
			_, err := r.errReader.buffer.Read(read)
			if err != nil {
				log.Println("err in out reading", err)
			} else {
				r.processError(strings.TrimSpace(string(read)))
			}
		}
	}
}

type stdoutWriter struct {
	buffer bytes.Buffer
}

func (w *stdoutWriter) Write(p []byte) (int, error) {
	return w.buffer.Write(p)
}

type stderrWriter struct {
	buffer bytes.Buffer
}

func (w *stderrWriter) Write(p []byte) (int, error) {
	return w.buffer.Write(p)
}

/*
	log.Println("Performing Http Post...")
	req := map[string]interface{}{"APIkey": "5rrtg3u564uiqr43fadf", "LogRequest": LogRequest{Date: "2021-06-17 22:26:43", Number: 12, Message: "hitext", Type: Important}}
	jsonReq, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:18769/api", "application/json;", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Println(bodyString)
*/

type LogRequest struct {
	Date    string
	Number  int
	Message string
	Type    Logtype
}

type Logtype string

const (
	Low       Logtype = "Low"
	Normal    Logtype = "Normal"
	Important Logtype = "Important"
	Error     Logtype = "Error"
)

type Activity struct {
	Date string
	Type Activitytype
}

type Activitytype string

const (
	Send              Activitytype = "Send"
	Recive            Activitytype = "Recive"
	Process           Activitytype = "Process"
	Backgroundprocess Activitytype = "Backgroundprocess"
)
