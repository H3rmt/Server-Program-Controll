package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

/*
struct representing a Program

can be started

contains
program (python/go/...)
File arg for programm, usually file but can be any series of args
APIKey to send data to Server
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
starts the program and readers
*/
func (pr *Program) Start() (err error) {
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

	err = cmd.Start()
	log.Println("Started Program: ", pr.Program, pr.Arguments)

	go func() {
		cmd.Wait()
		pr.stop = true
		pr.cmd = nil
		log.Println("Program finished: ", pr.Program, pr.Arguments)

		err := SendShutdown(pr)
		if err != nil {
			log.Println("err sending shutdown", err)
		}
		pr.logcounter = 0
	}()

	return
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
