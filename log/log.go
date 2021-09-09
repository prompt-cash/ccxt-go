package log

import (
	"errors"
	"github.com/spf13/viper"
	"math"
)

// A global variable so that log functions can be directly accessed
//var log Logger // TODO we don't use this. attach buffer and and writer to a new map with logger instance as key?

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

// Define a type for parameters in functions
type LoggerLevel string

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

type LoggerInstance int

const (
	InstanceZapLogger LoggerInstance = iota
	InstanceLogrusLogger
)

// expose the default logger in a variable.
// TODO add a config to change it at runtime? overkill
const DefaultLogger = InstanceZapLogger

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	// stdOut logging
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      LoggerLevel
	Color             bool

	// file logging
	EnableFile     bool
	FileJSONFormat bool
	FileLevel      LoggerLevel
	FileLocation   string

	// buffer logging (for HTTP admin API)
	EnableBuffer     bool
	BufferLines      int
	BufferJSONFormat bool
	BufferLevel      LoggerLevel

	// log adapter
	FilterPrefixes []string
}

// Creates a new logger config based on viper config settings.
func NewConfig(conf *viper.Viper) *Configuration {
	// we use viper as argument so that project's that don't use a global
	// viper config can create a temporary object for this call
	logLevel := NormalizeLogLevel(conf.GetString("Log.Level"))
	return &Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: conf.GetBool("Log.JSON"),
		Color:             conf.GetBool("Log.Color"),
		ConsoleLevel:      logLevel,

		EnableFile:     conf.GetBool("Log.EnableFile"),
		FileJSONFormat: conf.GetBool("Log.JSON"),
		FileLevel:      logLevel,
		FileLocation:   "log.log", // in working dir

		EnableBuffer:     conf.GetBool("Log.EnableBuffer"),
		BufferLines:      10000,
		BufferJSONFormat: true, // always use JSON for logs we server via API
		BufferLevel:      logLevel,

		FilterPrefixes: conf.GetStringSlice("Log.FilterPrefixes"),
	}
}

// NewLogger returns an instance of logger
func NewLogger(config *Configuration, loggerInstance LoggerInstance) (Logger, error) {
	var err error
	var logger Logger

	switch loggerInstance {
	case InstanceZapLogger:
		logger, err = newZapLogger(*config)
		if err != nil {
			return nil, err
		}
		break
		/*
			case InstanceLogrusLogger:
				logger, err = newLogrusLogger(*config)
				if err != nil {
					return nil, err
				}
				break
		*/

	default:
		return nil, errInvalidLoggerInstance
	}
	writer = Writer{
		Filter: config.FilterPrefixes,
		log:    logger,
	}
	return logger, nil
}

func NormalizeLogLevel(logLevel string) LoggerLevel {
	var nomalizedLogLevel LoggerLevel
	switch logLevel {
	case "info":
		nomalizedLogLevel = Info
	case "debug":
		nomalizedLogLevel = Debug
	case "warn":
		nomalizedLogLevel = Warn
	case "error":
		nomalizedLogLevel = Error
	case "fatal":
		nomalizedLogLevel = Fatal
	default:
		nomalizedLogLevel = Debug
	}
	return nomalizedLogLevel
}

// our log buffer (if enabled)
var buffer Buffer

// Get the buffer with recent log lines
func GetBuffer() *Buffer {
	return &buffer
}

func createBuffer(config Configuration) Buffer {
	return Buffer{
		MaxSize:   config.BufferLines,
		CleanSize: int(math.Floor(float64(config.BufferLines / 2.0))),
	}
}

var writer Writer

// Get the a struct that can be used as an io.Writer to write log messages to.
// Useful with packages that expect an io.Writer for their log, for example: http.Server
func GetWriter() *Writer {
	return &writer
}
