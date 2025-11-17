package common

import (
	"fmt"
	"log"
	"os"
	"time"
)

func NewLogger(prefic string) *log.Logger {
	return log.New(os.Stdout, prefic+"", log.Ldate|log.Ltime)
}

func LogOperation(operation string) {
	fmt.Printf("Operation: %s", time.Now().Format("2006-01-02 15:04:05")+operation)
}

func formatLogMessage(level, message string) string {
	return fmt.Sprintf("%s: %s", time.Now().Format("2006-01-02 15:04:05"), level, message)
}
