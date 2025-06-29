package singleton

import (
	"fmt"
	"sync"
	"time"
)

// Logger represents a singleton logger instance
type Logger struct {
	logLevel string
	mu       sync.Mutex // protects concurrent writes
}

var (
	loggerInstance *Logger
	loggerOnce     sync.Once
)

// GetLogger returns the singleton instance of Logger
func GetLogger() *Logger {
	loggerOnce.Do(func() {
		loggerInstance = &Logger{
			logLevel: "INFO",
		}
		fmt.Println("Creating new logger instance")
	})
	return loggerInstance
}

// Log writes a log message with timestamp
func (l *Logger) Log(level, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s: %s\n", timestamp, level, message)
}

// SetLogLevel sets the logging level
func (l *Logger) SetLogLevel(level string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
	fmt.Printf("Log level set to: %s\n", level)
}

// GetLogLevel returns the current log level
func (l *Logger) GetLogLevel() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.logLevel
}