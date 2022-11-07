package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "syarif").Info("Hello World")

	logger.WithField("username", "Hidayatulloh").
		WithField("name", "Syarif").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "syarif",
		"name":     "syarif Hidayatulloh",
	}).Infof("Hello world")
}
