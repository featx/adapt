package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type LoggerConfig struct {
	Name string `json:"name" yaml:"name" toml:"name" xml:"name"`
	Output string `json:"output" yaml:"output" toml:"output" xml:"output"`
	Prefix string `json:"prefix" yaml:"prefix" toml:"prefix" xml:"prefix"`
	Level string `json:"level" yaml:"level" toml:"level" xml:"level"`
	Pattern string `json:"pattern" yaml:"pattern" toml:"pattern" xml:"pattern"`
}

type LoggerOutputType string

const (
	OutputFile LoggerOutputType = "file"
)

func ConfigLogger(logConfig LoggerConfig) *log.Logger {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
	if strings.HasPrefix(logConfig.Output, string(OutputFile)) {
		file, err := os.Create(logConfig.Output[len(OutputFile) + 2:])
		if err != nil {
			panic(fmt.Errorf("construct logger from loggerConfig error: %s", err.Error()))
		}
		return log.New(file, logConfig.Prefix, flag)
	}
	return log.New(os.Stdout, logConfig.Prefix, flag)
}

func MapLoggerConfigs(logConfigs []LoggerConfig) map[string]LoggerConfig {
	if logConfigs == nil {
		return nil
	}
	result := make(map[string]LoggerConfig)
	for _, logConfig := range logConfigs {
		result[logConfig.Name] = logConfig
	}
	return result
}