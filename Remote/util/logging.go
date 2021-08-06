package util

import (
	"fmt"
	lg "log"
)

var stretch = "9"

func Log(pefx string, message ...interface{}) {
	prn := fmt.Sprintf("%-"+stretch+"s|", pefx)

	for _, mess := range message {
		prn += fmt.Sprintf("%v", mess)
	}
	lg.Println(prn)
}
