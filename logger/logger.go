package logger

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

func init() {
	showNumberLine, _ := beego.AppConfig.Bool("logger.numberLine")
	levelString := beego.AppConfig.String("logger.level")
	logModels := beego.AppConfig.Strings("logger.logger")
	for _, logModel := range logModels {
		if logModel == "console" {
			beego.SetLogger("console", "")
		}
		if logModel == "file" {
			filename := beego.AppConfig.String("logger.filename")
			config := fmt.Sprintf(`{"filename":"%s","maxdays":30}`, filename)
			beego.SetLogger("file", config)
		}
	}
	if showNumberLine {
		beego.SetLogFuncCall(showNumberLine)
	}
	level, _ := Levelify(levelString)
	beego.SetLevel(level)
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

var levels = map[string]int{
	"DEBUG":    logs.LevelDebug,
	"INFO":     logs.LevelInformational,
	"WARN":     logs.LevelWarning,
	"ERROR":    logs.LevelError,
	"CRITICAL": logs.LevelCritical,
}

var levelKeys = []string{"DEBUG", "INFO", "WARN", "ERROR", "CRITICAL"}
