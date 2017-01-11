package log

import (
	"encoding/json"
	"log"
)

// Log types
var (
	typeError = "error"
)

// JSON parses the input JSON data and prints it.
func JSON(logType string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		Error(err)
		return
	}
	Log(logType, string(b))
}

// Error prints the input error.
func Error(err error) {
	Log(typeError, err.Error())
}

// Log prints the input string.
func Log(logType, s string) {
	log.Printf("[%s] %s\n", logType, s)
}
