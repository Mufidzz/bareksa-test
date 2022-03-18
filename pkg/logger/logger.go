package logger

import (
	"log"
)

func Error(err error) {
	// Add Code here to Advanced Logger ex. send to monitoring service, customize logging format, etc
	log.Printf("[ERROR] %v", err)
}
