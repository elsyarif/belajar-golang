package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLoggert(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello World")
}
