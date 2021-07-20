package main

import (
	"bytes"
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
}

/*
starts the program and readers
*/
func (pr *Program) Start() {
	cmd := exec.Command(pr.Program, pr.Arguments...)

	pr.reader = Reader{cmdrunning: false}

	cmd.Stdout = &pr.reader.outReader
	cmd.Stderr = &pr.reader.errReader

	go pr.reader.process()
	cmd.Wait()
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

/*
Custom Reader containing out and err Writer
*/
type Reader struct {
	cmdrunning bool
	outReader  stdoutWriter
	errReader  stderrWriter
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
	for r.cmdrunning {
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
