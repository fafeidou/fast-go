package logging

import (
	"fast-go/conf"
	"fmt"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {

	return fmt.Sprintf("%s%s", conf.App.RuntimeRootPath, conf.App.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		conf.App.LogSaveName,
		time.Now().Format(conf.App.TimeFormat),
		conf.App.LogFileExt,
	)
}
