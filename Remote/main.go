package main

import (
	"fmt"
	"log"
	"os"
)

// list of all programs
var programs []Program

/*
main method with arguments = programfile APIKey programfile APIKey ...
*/
func main() {
	// Create programs from args
	args := os.Args[1:]
	if len(args) == 0 || len(args)%2 != 0 {
		// must be even number of args
		log.Println("invalid args: programfile APIKey programfile APIKey ...")
		return
	}
	for i := 0; i < len(args); i += 2 {
		// return is program already added
		for _, v := range programs {
			if v.File == args[i] {
				log.Println("duplicate file:", args[i])
				return
			}
		}
		// check if file exists
		if _, err := os.Stat(args[i]); err == nil {
			programs = append(programs, Program{APIKey: args[i+1], File: args[i]})
		} else {
			log.Println("invalid program:", args[i])
			return
		}
	}
	for _, v := range programs {
		log.Println(fmt.Sprintf("%s Key:%s", v.File, v.APIKey))
	}
}
