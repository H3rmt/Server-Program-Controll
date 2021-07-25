package api

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// list of all programs
var Programs []Program

/*
struct representing a Program

contains
program  (python/go/...)
Arguments  arg for programm, usually file but can be any series of args
APIKey  to send data to Server
reader  containing out and err reader
stop  bool if program is running
cmd  reference to Command
logcounter  counts logs for out and err reader
*/
type Program struct {
	Program    string
	Arguments  []string
	APIKey     string
	reader     Reader
	stop       bool
	cmd        *exec.Cmd
	logcounter int
}

/*
finds the corresponding program from list of programs
*/
func getProgramm_IDfromAPIKey(APIKey string) (*Program, error) {
	for i := 0; i < len(Programs); i++ {
		if Programs[i].APIKey == APIKey {
			log.Println("Program found:", APIKey, Programs[i].Program, Programs[i].Arguments)
			return &Programs[i], nil
		}
	}
	return &Program{}, &InvalidAPIKeyerror{}
}

/*
starts the program and readers
*/
func (pr *Program) Start() error {
	if pr.cmd != nil {
		return fmt.Errorf("program running")
	}
	cmd := exec.Command(pr.Program, pr.Arguments...)
	pr.cmd = cmd

	pr.reader = Reader{programparent: pr}

	cmd.Stdout = &pr.reader.outReader
	cmd.Stderr = &pr.reader.errReader

	pr.stop = false
	go pr.reader.process()

	pr.logcounter = 0

	err := cmd.Start()
	if err != nil {
		log.Println("Error Starting Program: ", pr.Program, pr.Arguments)
		return err
	} else {
		log.Println("Started Program: ", pr.Program, pr.Arguments)
	}

	go func() {
		// Delay this Message so it doesnt mix up with start response
		time.Sleep(30 * time.Millisecond)
		err = SendStateChange(pr, true)
		if err != nil {
			log.Println("err sending start", err)
		}
	}()

	go func() {
		cmd.Wait()
		pr.stop = true
		pr.cmd = nil
		log.Println("Program finished: ", pr.Program, pr.Arguments)

		err := SendStateChange(pr, false)
		if err != nil {
			log.Println("err sending shutdown", err)
		}
	}()

	return nil
}

/*
stops the program
*/
func (pr *Program) Stop() (err error) {
	if pr.cmd != nil {
		err = pr.cmd.Process.Kill()
		if err != nil {
			pr.cmd = nil
		}
	} else {
		err = fmt.Errorf("program was not started")
	}
	return
}

/*
finds logtype of logmessage and returns message without prefix
*/
func processOutLevel(message string) (string, Logtype) {
	if strings.HasPrefix(message, "LOW|") {
		return strings.Replace(message, "LOW|", "", 1), Logtype(Low)
	} else if strings.HasPrefix(message, "NORMAL|") {
		return strings.Replace(message, "NORMAL|", "", 1), Logtype(Normal)
	} else if strings.HasPrefix(message, "IMPORTANT|") {
		return strings.Replace(message, "IMPORTANT|", "", 1), Logtype(Important)
	} else {
		return message, Logtype(Normal)
	}
}

/*
Custom Reader containing out and err Writer
*/
type Reader struct {
	programparent *Program
	outReader     stdoutWriter
	errReader     stderrWriter
}

/*
process stdout Info
*/
func (r *Reader) processOutput(out string) {
	newout, logtype := processOutLevel(string(out))
	log.Println(r.programparent.Program, r.programparent.Arguments, "out:", newout)
	err := SendLog(newout, r.programparent, logtype)
	if err != nil {
		log.Println("err sending out", err)
	}
	r.programparent.logcounter++
}

/*
process stderr Info
*/
func (r *Reader) processError(err string) {
	log.Println(r.programparent.Program, r.programparent.Arguments, "err:", string(err))
	errr := SendLog(string(err), r.programparent, Logtype(Error))
	if errr != nil {
		log.Println("err sending err", errr)
	}
	r.programparent.logcounter++
}

/*
runs parallel to Program and reads stdout and stderr
feeds bot in processOutput and processError
*/
func (r *Reader) process() {
	for !r.programparent.stop {
		time.Sleep(3 * time.Millisecond)
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
