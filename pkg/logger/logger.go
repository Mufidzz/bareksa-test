package logger

import (
	"fmt"
	"log"
)

func Errorf(format string, parameters ...interface{}) error {
	// Add Code here to Advanced Logger ex. send to monitoring service, customize logging format, etc
	log.Printf("[ERROR] "+format, parameters...)

	return fmt.Errorf(format, parameters...)
}
