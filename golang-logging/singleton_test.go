package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello Info")
	logrus.Warn("Hello Warn")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Hello Info")
	logrus.Warn("Hello Warn")
}
