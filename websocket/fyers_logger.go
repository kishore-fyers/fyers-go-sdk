package fyersgosdk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// FyersLogger represents a logger instance for Fyers API
type FyersLogger struct {
	service    string
	level      string
	stackLevel int
	logger     *log.Logger
	logFile    *os.File
	location   string
}

// NewFyersLogger creates a new FyersLogger instance
func NewFyersLogger(service, level string, stackLevel int, logPath string) *FyersLogger {
	if stackLevel == 0 {
		stackLevel = 4
	}

	// Create log directory if it doesn't exist
	if logPath != "" {
		err := os.MkdirAll(logPath, 0755)
		if err != nil {
			log.Printf("Failed to create log directory: %v", err)
		}
	}

	// Create log file
	var logFile *os.File
	var err error
	if logPath != "" {
		// Determine log filename based on service
		var logFileName string
		if strings.Contains(service, "Order") {
			logFileName = "fyersOrderSocket.log"
		} else {
			logFileName = "fyersDataSocket.log"
		}

		logFilePath := filepath.Join(logPath, logFileName)
		logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Failed to open log file: %v", err)
		}
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	if logFile != nil {
		logger.SetOutput(logFile)
	}

	return &FyersLogger{
		service:    service,
		level:      level,
		stackLevel: stackLevel,
		logger:     logger,
		logFile:    logFile,
	}
}

// getLocation returns the function name, line number, and module name
func (f *FyersLogger) getLocation() string {
	if f.location != "" {
		return f.location
	}

	pc, file, line, ok := runtime.Caller(f.stackLevel)
	if !ok {
		return "[unknown:0] unknown"
	}

	funcName := runtime.FuncForPC(pc).Name()
	// Extract just the function name without package path
	parts := strings.Split(funcName, ".")
	funcName = parts[len(parts)-1]

	// Extract just the filename without path
	parts = strings.Split(file, "/")
	fileName := parts[len(parts)-1]

	f.location = fmt.Sprintf("[%s:%d] %s", funcName, line, fileName)
	return f.location
}

// formatMessage formats the log message with timestamp and location
func (f *FyersLogger) formatMessage(level, msg string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	location := f.getLocation()
	return fmt.Sprintf("[%s] [%s] [%s] %s - %s", timestamp, f.service, level, location, msg)
}

// Error logs an error message
func (f *FyersLogger) Error(msg interface{}) {
	message := fmt.Sprintf("%v", msg)
	formattedMsg := f.formatMessage("ERROR", message)
	f.logger.Println(formattedMsg)
}

// Info logs an info message
func (f *FyersLogger) Info(msg interface{}) {
	message := fmt.Sprintf("%v", msg)
	formattedMsg := f.formatMessage("INFO", message)
	f.logger.Println(formattedMsg)
}

// Debug logs a debug message
func (f *FyersLogger) Debug(msg interface{}) {
	if f.level == "DEBUG" {
		message := fmt.Sprintf("%v", msg)
		formattedMsg := f.formatMessage("DEBUG", message)
		f.logger.Println(formattedMsg)
	}
}

// Exception logs an exception message (equivalent to Python's exception method)
func (f *FyersLogger) Exception(msg interface{}) {
	message := fmt.Sprintf("%v", msg)
	formattedMsg := f.formatMessage("EXCEPTION", message)
	f.logger.Println(formattedMsg)
}

// Close closes the log file
func (f *FyersLogger) Close() {
	if f.logFile != nil {
		f.logFile.Close()
	}
}
