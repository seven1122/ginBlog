package logging

import (
	"fmt"
	"seven1122/ginBlog/pkg/setting"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RunTimeRootPath, setting.AppSetting.LogSavePath)

}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)

}
