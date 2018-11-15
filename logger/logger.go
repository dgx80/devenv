package logger

import "log"

var Verbose bool

//Info log only on verbose
func Info(message string) {
	log.Println(message)
}
