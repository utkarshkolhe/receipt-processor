package logger

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "sync"
    "path/filepath"
    "utkarsh/Fetch/api/config"
)

type Logger struct {
  fileLogger    *log.Logger
  consoleLogger *log.Logger
}

var (
    Instance *Logger
    once     sync.Once
)

func init() {
    Instance = Init(config.LogFile)
}

// Initialize the logger and create or open the specified log file.
func Init(logFilePath string) *Logger {
    once.Do(func() {
        file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err != nil {
            log.Fatalf(config.ErrLogFileFailed+": %v", err)
        }
        // File logger without date and time
        fileLogger := log.New(file, "", log.Ldate|log.Ltime) // No flags for custom formatting

        // Console logger with date and time
        consoleLogger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

        Instance = &Logger{
            fileLogger:    fileLogger,
            consoleLogger: consoleLogger,
        }

    })
    return Instance
}

// getCallerInfo retrieves the file name and line number of the caller.
func getCallerInfo() string {
    _, file, line, ok := runtime.Caller(2) // 2 levels up from where getCallerInfo is called
    if !ok {
        return "unknown:0"
    }
    return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

// Info logs an informational message with caller info.
func (l *Logger) Info(message string) {
    if config.FileLog{
      l.fileLogger.Printf("INFO: %s - %s", getCallerInfo(), message)
    }
    if config.ConsoleLog{
        l.consoleLogger.Printf("INFO: %s - %s", getCallerInfo(), message)
    }

}

// Warn logs a warning message with caller info.
func (l *Logger) Warn(message string) {
    if config.FileLog{
      l.fileLogger.Printf("WARN: %s - %s", getCallerInfo(), message)
    }
    if config.ConsoleLog{
      l.consoleLogger.Printf("WARN: %s - %s", getCallerInfo(), message)
    }
}

// Error logs an error message with caller info.
func (l *Logger) Error(message string) {
    if config.FileLog{
      l.fileLogger.Printf("ERROR: %s - %s", getCallerInfo(), message)
    }
    if config.FileLog{
      l.consoleLogger.Printf("ERROR: %s - %s", getCallerInfo(), message)
    }
}
