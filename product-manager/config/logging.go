package config

import (
	"io"
	"log"
	"os"
)

type LogConfig struct {
	LogFile string
}

func NewLogConfig(logFile string) *LogConfig {
	return &LogConfig{
		LogFile: logFile,
	}
}

func (l *LogConfig) LoggingSettings() {
	logfile, err := os.OpenFile(l.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
