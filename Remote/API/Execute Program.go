package api

import (
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
cmd  reference to Command
logcounter  counts logs for out and err reader
*/
type Program struct {
	Name          string
	Program       string
	Arguments     []string
	APIKey        string
	reader        Reader
	cmd           *exec.Cmd
	logcounter    int
	LogasActivity bool
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

	pr.reader = Reader{}

	pr.reader.outReader.programparent = pr
	pr.reader.errReader.programparent = pr

	cmd.Stdout = &pr.reader.outReader
	cmd.Stderr = &pr.reader.errReader

	pr.logcounter = 0

	err := cmd.Start()
	if err != nil {
		log.Println("Error Starting Program: ", pr.Name)
		return err
	} else {
		log.Println("Started Program: ", pr.Name)
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
		pr.cmd = nil
		log.Println("Program finished: ", pr.Name)

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
checks if return is log or activity
*/
func CheckLog(message string) string {
	if strings.HasPrefix(message, "LOW|") || strings.HasPrefix(message, "NORMAL|") || strings.HasPrefix(message, "IMPORTANT|") {
		return "Log"
	} else if strings.HasPrefix(message, "[Send]") || strings.HasPrefix(message, "[Recive]") || strings.HasPrefix(message, "[Process]") || strings.HasPrefix(message, "[Backgroundprocess]") {
		remove := strings.TrimSpace(strings.Replace(message, strings.SplitN(message, "]", 2)[0]+"]", "", 1))
		if strings.HasPrefix(remove, "LOW|") || strings.HasPrefix(remove, "NORMAL|") || strings.HasPrefix(remove, "IMPORTANT|") {
			return "Activity Log"
		} else {
			return "Activity"
		}
	}
	return ""
}

/*
finds logtype of logmessage and returns message without prefix
*/
func processLogLevel(message string) (string, Logtype) {
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
finds logtype of logmessage and returns message without prefix
*/
func processActivityLevel(message string) Activitytype {
	if strings.HasPrefix(message, "[Backgroundprocess]") {
		return Activitytype(Backgroundprocess)
	} else if strings.HasPrefix(message, "[Process]") {
		return Activitytype(Process)
	} else if strings.HasPrefix(message, "[Recive]") {
		return Activitytype(Recive)
	} else if strings.HasPrefix(message, "[Send]") {
		return Activitytype(Send)
	} else {
		return Activitytype(Process)
	}
}

/*
Custom Reader containing out and err Writer
*/
type Reader struct {
	outReader stdoutWriter
	errReader stderrWriter
}

type stdoutWriter struct {
	programparent *Program
}

func (w *stdoutWriter) Write(p []byte) (int, error) {
	w.processOutput(strings.TrimSpace(string(p)))
	return len(p), nil
}

/*
process stdout Info
*/
func (w *stdoutWriter) processOutput(out string) {
	log.Println(w.programparent.Name, "out:", out)
	outtype := CheckLog(string(out))
	if outtype == "Log" {
		newout, logtype := processLogLevel(strings.TrimSpace(string(out)))
		err := SendLog(strings.TrimSpace(newout), w.programparent, logtype)
		if err != nil {
			log.Println("err sending out log", err)
		}
	} else if outtype == "Activity" {
		activitytype := processActivityLevel(strings.TrimSpace(string(out)))
		err := SendActivity(w.programparent, activitytype)
		if err != nil {
			log.Println("err sending out activity", err)
		}
	} else if outtype == "Activity Log" {
		activitytype := processActivityLevel(strings.TrimSpace(string(out)))
		err := SendActivity(w.programparent, activitytype)
		if err != nil {
			log.Println("err sending out activity", err)
		}
		removed := strings.TrimSpace(strings.Replace(string(out), strings.SplitN(string(out), "]", 2)[0]+"]", "", 1))
		newout, logtype := processLogLevel(strings.TrimSpace(removed))
		err = SendLog(strings.TrimSpace(newout), w.programparent, logtype)
		if err != nil {
			log.Println("err sending out log", err)
		}
	}
}

type stderrWriter struct {
	programparent *Program
}

func (w *stderrWriter) Write(p []byte) (int, error) {
	str := string(p)
	if strings.Contains(str, "\n") {
		for _, line := range strings.Split(str, "\n") {
			w.processError(strings.TrimSpace(line))
		}
	} else {
		w.processError(strings.TrimSpace(str))
	}
	return len(p), nil
}

/*
process stderr Info
*/
func (w *stderrWriter) processError(err string) {
	log.Println(w.programparent.Name, "err:", string(err))
	errr := SendLog(strings.TrimSpace(string(err)), w.programparent, Logtype(Error))
	if errr != nil {
		log.Println("err sending err", errr)
	}
	w.programparent.logcounter++
}
