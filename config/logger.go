package config

import (
	"github.com/sirupsen/logrus"
)

// NewLogger is
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	return logger
}
