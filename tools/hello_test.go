package tools_test

import (
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"imjuni.logrustest.com/tools"
)

func TestHello(t *testing.T) {
	test.NewGlobal()

	// log message not displayed
	tools.Hello("world")
}
