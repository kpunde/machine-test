package logger

import "testing"

func TestInitService(t *testing.T) {
	l := InitService("../../logs/testLog.log", 4)
	l.Info("Test message")
}
