package tools

import (
	"fmt"

	"log"

	"github.com/sirupsen/logrus"
)

// Hello hello function for testing
func Hello(who string) string {
	logrus.Debug("I want disply!! plz")

	// Plz help, I want same behavior
	log.Printf("This log message displayed!")

	return fmt.Sprintf("hello %s!", who)
}
