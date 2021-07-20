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
	Program   string
	Arguments []string
	APIKey    string
	reader    Reader
	cmd       *exec.Cmd
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

	pr.reader = Reader{stop: false}

	cmd.Stdout = &pr.reader.outReader
	cmd.Stderr = &pr.reader.errReader

	go pr.reader.process()

	err = cmd.Start()
	log.Println("Started Program: ", pr.Program, pr.Arguments)

	go func() {
		cmd.Wait()
		pr.reader.stop = true
		pr.cmd = nil
		log.Println("Program finished: ", pr.Program, pr.Arguments)
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

/*
Custom Reader containing out and err Writer
*/
type Reader struct {
	stop      bool
	outReader stdoutWriter
	errReader stderrWriter
}

/*
process stdout Info
*/
func (r *Reader) processOutput(out string) {
	log.Println("out:", string(out))
}

/*
process stderr Info
*/
func (r *Reader) processError(err string) {
	log.Println("err:", string(err))
}

/*
runs parallel to Program and reads stdout and stderr
feeds bot in processOutput and processError
*/
func (r *Reader) process() {
	for !r.stop {
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
