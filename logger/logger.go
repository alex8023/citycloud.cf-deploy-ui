package logger

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"strings"
	"fmt"
)

var log *logs.BeeLogger

var levels = map[string]int{
	"DEBUG": logs.LevelDebug,
	"INFO":  logs.LevelInformational,
	"WARN":  logs.LevelWarning,
	"ERROR": logs.LevelError,
	"CRITICAL":  logs.LevelCritical,
}

var levelKeys = []string{"DEBUG", "INFO", "WARN", "ERROR", "CRITICAL"}

func init() {
	log = logs.NewLogger(10000)
	
	showNumberLine,err := beego.AppConfig.Bool("logger.numberLine")

	if err == nil{
		if showNumberLine {
			log.EnableFuncCallDepth(true)
			log.SetLogFuncCallDepth(3)
		}
	}
	
	levelString := beego.AppConfig.String("logger.level")
	
	level,errs := Levelify(levelString)
	if errs != nil {
		fmt.Println(errs)
	}
	log.SetLevel(level)
	
	logModels := beego.AppConfig.Strings("logger.logger")
	
	for _,logModel := range logModels {
		if logModel == "console" {
			log.SetLogger("console","")
		}
		
		if logModel  == "file" {
			filename := beego.AppConfig.String("logger.filename")
			if filename =="" {
				filename = "logs.log"
			}
			log.SetLogger("file", `{
				"filename": "`+filename+`",
				"maxdays": 365
				}`)
		}
	}
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

func Debug(format string, v ...interface{}){
	log.Debug(format,v)
}

func Info(format string, v ...interface{}){
	log.Trace(format,v)
}

func Warn(format string, v ...interface{}){
	log.Warn(format,v)
}

func Error(format string, v ...interface{}){
	log.Error(format,v)
}

func Critical(format string, v ...interface{}){
	log.Critical(format,v)
}