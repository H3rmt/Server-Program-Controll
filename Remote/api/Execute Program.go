package api

import (
	"fmt"
	"os/exec"
	"strings"

	"Remote/util"
)

// list of all programs
var Programs []Program

/*
struct representing a Program

contains
program  (python/go/...)
Arguments  arg for programm, usually file but can be any series of args
APIKey  to send data to Server
Dir       directory where to execute cmd
reader  containing out and err reader
cmd  reference to Command
*/
type Program struct {
	Name      string
	Program   string
	Arguments []string
	APIKey    string
	Dir       string
	reader    Reader
	cmd       *exec.Cmd
	running   bool
	// LogasActivity bool
}

/*
finds the corresponding program from list of programs
*/
func getprogrammIdfromapikey(APIKey string) (*Program, error) {
	for i := 0; i < len(Programs); i++ {
		if Programs[i].APIKey == APIKey {
			util.Log("EXEC PR", "Program found:", APIKey, Programs[i].Program, Programs[i].Arguments)
			return &Programs[i], nil
		}
	}
	return &Program{}, &InvalidAPIKeyerror{}
}

// Start
func (pr *Program) Start() error {
	if pr.running {
		return fmt.Errorf("program already running")
	}
	pr.cmd = exec.Command(pr.Program, pr.Arguments...)
	pr.cmd.Dir = pr.Dir

	pr.reader = Reader{}

	pr.reader.outReader.programparent = pr
	pr.reader.errReader.programparent = pr

	pr.cmd.Stdout = &pr.reader.outReader
	pr.cmd.Stderr = &pr.reader.errReader

	err := SendStateChange(pr, true)
	if err != nil {
		util.Log("EXEC PR", "error sending start", err)
	}
	pr.running = false

	err = pr.cmd.Start()
	if err != nil {
		util.Log("EXEC PR", "error Starting Program: ", pr.Name)
		return err
	} else {
		util.Log("EXEC PR", "Started Program: ", pr.Name)
	}

	go func() {
		err := pr.cmd.Wait()
		pr.running = false
		util.Log("EXEC PR", "Program finished: ", pr.Name, "Error:", err)
		if err != nil {
			err := SendLog(err.Error(), pr, Important)
			if err != nil {
				util.Log("EXEC PR", "err sending error Log", err)
			}
		}

		err = SendStateChange(pr, false)
		if err != nil {
			util.Log("EXEC PR", "err sending shutdown", err)
		}
	}()

	return nil
}

// Stop /*
func (pr *Program) Stop() (err error) {
	if pr.running {
		err = pr.cmd.Process.Kill()
		pr.running = false
	} else {
		err = fmt.Errorf("program not running")
	}
	return
}

// CheckLog /*
func CheckLog(message string) string {
	if strings.HasPrefix(message, "LOW|") || strings.HasPrefix(message, "NORMAL|") || strings.HasPrefix(message, "IMPORTANT|") {
		return "Log"
	} else if strings.HasPrefix(message, "[Send]") || strings.HasPrefix(message, "[Receive]") || strings.HasPrefix(message, "[Process]") || strings.HasPrefix(message, "[Backgroundprocess]") {
		remove := strings.TrimSpace(strings.Replace(message, strings.SplitN(message, "]", 2)[0]+"]", "", 1))
		if strings.HasPrefix(remove, "LOW|") || strings.HasPrefix(remove, "NORMAL|") || strings.HasPrefix(remove, "IMPORTANT|") {
			return "Activity Log"
		} else {
			return "Activity"
		}
	}
	return "Log" // default messages interpreted as log
}

/*
finds logtype of logmessage and returns message without prefix
*/
func processLogLevel(message string) (string, Logtype) {
	if strings.HasPrefix(message, "LOW|") {
		return strings.Replace(message, "LOW|", "", 1), Low
	} else if strings.HasPrefix(message, "NORMAL|") {
		return strings.Replace(message, "NORMAL|", "", 1), Normal
	} else if strings.HasPrefix(message, "IMPORTANT|") {
		return strings.Replace(message, "IMPORTANT|", "", 1), Important
	} else {
		return message, Logtype(Normal)
	}
}

/*
finds logtype of logmessage and returns message without prefix
*/
func processActivityLevel(message string) Activitytype {
	if strings.HasPrefix(message, "[Backgroundprocess]") {
		return Backgroundprocess
	} else if strings.HasPrefix(message, "[Process]") {
		return Process
	} else if strings.HasPrefix(message, "[Receive]") {
		return Recive
	} else if strings.HasPrefix(message, "[Send]") {
		return Send
	} else {
		return Process
	}
}

// Reader /*
type Reader struct {
	outReader stdoutWriter
	errReader stderrWriter
}

type stdoutWriter struct {
	programparent *Program
}

func (w *stdoutWriter) Write(p []byte) (int, error) {
	str := strings.TrimSpace(string(p))
	if len(str) == 0 {
		return len(p), nil
	}
	if strings.Contains(str, "\n") {
		util.Log("split")
		for _, line := range strings.Split(str, "\n") {
			str := strings.TrimSpace(line)
			if len(str) == 0 {
				continue
			}
			w.processOutput(str)
		}
	} else {
		w.processOutput(strings.TrimSpace(str))
	}
	return len(p), nil
}

/*
process stdout Info
*/
func (w *stdoutWriter) processOutput(out string) {
	util.Log("EXEC PR", w.programparent.Name, " out: ", ">>", out, "<<")
	outtype := CheckLog(string(out))
	if outtype == "Log" {
		newout, logtype := processLogLevel(strings.TrimSpace(out))
		err := SendLog(strings.TrimSpace(newout), w.programparent, logtype)
		if err != nil {
			util.Log("EXEC PR", "err sending out log", err)
		}
	} else if outtype == "Activity" {
		activitytype := processActivityLevel(strings.TrimSpace(out))
		err := SendActivity(w.programparent, activitytype)
		if err != nil {
			util.Log("EXEC PR", "err sending out activity", err)
		}
	} else if outtype == "Activity Log" {
		activitytype := processActivityLevel(strings.TrimSpace(out))
		err := SendActivity(w.programparent, activitytype)
		if err != nil {
			util.Log("EXEC PR", "err sending out activity", err)
		}
		removed := strings.TrimSpace(strings.Replace(string(out), strings.SplitN(out, "]", 2)[0]+"]", "", 1))
		newout, logtype := processLogLevel(strings.TrimSpace(removed))
		err = SendLog(strings.TrimSpace(newout), w.programparent, logtype)
		if err != nil {
			util.Log("EXEC PR", "err sending out log", err)
		}
	}
}

type stderrWriter struct {
	programparent *Program
}

func (w *stderrWriter) Write(p []byte) (int, error) {
	w.processError(strings.TrimSpace(string(p)))
	return len(p), nil
}

/*
process stderr Info
*/
func (w *stderrWriter) processError(err string) {
	util.Log("EXEC PR", w.programparent.Name, "err:", err)
	errr := SendLog(strings.TrimSpace(err), w.programparent, Error)
	if errr != nil {
		util.Log("EXEC PR", "err sending err", errr)
	}
}
