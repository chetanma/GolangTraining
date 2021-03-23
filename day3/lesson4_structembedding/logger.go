package main

import "fmt"

type Logger struct {
	logs []string
}

func NewLogger() Logger {
	return Logger{
		logs: make([]string, 0, 1000),
	}
}

func (log *Logger) Info(info string) {
	logItem := "INFO: " + info
	log.logs = append(log.logs, logItem)
}

func (log *Logger) Debug(debug string) {
	logItem := "DEBUG: " + debug
	log.logs = append(log.logs, logItem)
}

func (log *Logger) DumpToConsole() {
	for _, l := range log.logs {
		fmt.Println(l)
	}
}
