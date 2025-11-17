package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	writer *bufio.Writer
	file   *os.File
	mu     sync.Mutex
}

func NewLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	l := &Logger{writer: bufio.NewWriter(file), file: file}
	return l, nil
}

func (logger *Logger) Log(msg string) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logger.writer.WriteString(fmt.Sprintf("[%s] %s\n", msg, timestamp))
	logger.writer.Flush()
}

func (logger *Logger) Close() error {
	logger.writer.Flush()
	return logger.file.Close()
}

func main() {
	logger, _ := NewLogger("test.log")
	defer logger.Close()
	logger.Log("Application started")
	logger.Log("Processing data...")
}
