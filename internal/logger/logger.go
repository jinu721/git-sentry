package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	logFile *os.File
}

func NewLogger(gitsentryDir string) (*Logger, error) {
	logPath := filepath.Join(gitsentryDir, "logs", "gitsentry.log")
	
	// Create logs directory if it doesn't exist
	logsDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, err
	}
	
	// Open log file in append mode
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	
	return &Logger{logFile: logFile}, nil
}

func (l *Logger) Log(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)
	l.logFile.WriteString(logLine)
}

func (l *Logger) Info(message string) {
	l.Log("INFO", message)
}

func (l *Logger) Warn(message string) {
	l.Log("WARN", message)
}

func (l *Logger) Error(message string) {
	l.Log("ERROR", message)
}

func (l *Logger) Close() {
	if l.logFile != nil {
		l.logFile.Close()
	}
}