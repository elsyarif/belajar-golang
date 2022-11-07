package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormater(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Test Formater writer")
	logger.Warn("Test Formater writer")
	logger.Error("Test Formater writer")
}
