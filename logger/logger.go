package logger

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

var log *logs.BeeLogger

var levels = map[string]int{
	"DEBUG":    logs.LevelDebug,
	"INFO":     logs.LevelInformational,
	"WARN":     logs.LevelWarning,
	"ERROR":    logs.LevelError,
	"CRITICAL": logs.LevelCritical,
}

var levelKeys = []string{"DEBUG", "INFO", "WARN", "ERROR", "CRITICAL"}

func init() {
	showNumberLine, _ := beego.AppConfig.Bool("logger.numberLine")
	levelString := beego.AppConfig.String("logger.level")
	logConfig := make(map[string]LoggerModel)
	logModels := beego.AppConfig.Strings("logger.logger")
	for _, logModel := range logModels {
		if logModel == "console" {
			logConfig["console"] = LoggerModel{
				Name:   "console",
				Config: "",
			}
		}
		if logModel == "file" {
			filename := beego.AppConfig.String("logger.filename")
			logConfig["file"] = LoggerModel{
				Name: "file",
				Config: `{
				"filename": "` + filename + `",
				"maxdays": 365
				}`,
			}
		}
	}
	log = NewAppLogger(levelString, logConfig, showNumberLine).Logger
}

type LoggerModel struct {
	Name   string
	Config string
}

type AppLogger struct {
	Logger          *logs.BeeLogger
	Level           string
	LoggerConfig    map[string]LoggerModel
	LogerNumberLine bool
}

func NewAppLogger(level string, loggerConfig map[string]LoggerModel, loggerNumberLine bool) (appLoger AppLogger) {
	appLoger.Logger = logs.NewLogger(10000)
	appLoger.Level = strings.ToUpper(level)
	appLoger.LoggerConfig = loggerConfig
	appLoger.LogerNumberLine = loggerNumberLine

	lev, err := Levelify(level)

	if err != nil {
		fmt.Println(err)
	}
	appLoger.Logger.SetLevel(lev)

	if loggerNumberLine {
		appLoger.Logger.EnableFuncCallDepth(true)
		appLoger.Logger.SetLogFuncCallDepth(3)
	}

	for key, v := range loggerConfig {
		if key == "file" || key == "console" || key == "conn" || key == "stmp" {
			appLoger.Logger.SetLogger(v.Name, v.Config)
		}

	}
	return
}

func Levelify(levelString string) (int, error) {
	upperLevelString := strings.ToUpper(levelString)
	level, ok := levels[upperLevelString]
	if !ok {
		expected := strings.Join(levelKeys, ", ")
		return logs.LevelError, fmt.Errorf("Unknown LogLevel string '%s', expected one of [%s]", levelString, expected)
	}
	return level, nil
}

func Debug(format string, v ...interface{}) {
	log.Debug(format, v)
}

func Info(format string, v ...interface{}) {
	log.Trace(format, v)
}

func Warn(format string, v ...interface{}) {
	log.Warn(format, v)
}

func Error(format string, v ...interface{}) {
	log.Error(format, v)
}

func Critical(format string, v ...interface{}) {
	log.Critical(format, v)
}
