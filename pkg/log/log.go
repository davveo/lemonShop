package log

import (
	"fmt"
	"github.com/davveo/lemonShop/conf"
	"github.com/davveo/lemonShop/pkg/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Init() error {
	var (
		err    error
		logCfg = conf.Conf.Log
	)

	filePath := fmt.Sprintf("%s%s", logCfg.RuntimeRootPath, logCfg.LogSavePath)
	fileName := fmt.Sprintf("%s%s.%s",
		logCfg.LogSaveName,
		time.Now().Format(logCfg.TimeFormat),
		logCfg.LogFileExt,
	)
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
		return err
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
	return nil
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
