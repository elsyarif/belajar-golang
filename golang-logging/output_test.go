package golang_logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Test Output writer")
	logger.Warn("Test Output writer")
	logger.Error("Test Output writer")
}
