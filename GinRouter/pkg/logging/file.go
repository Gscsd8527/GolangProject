package logging

import (
	"GinRouter/pkg/setting"
	"fmt"
)

//获取日志文件路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}


//获取日志文件名字
func getLogFileName() string {
	//return fmt.Sprintf("%s%s.%s",
	//	setting.AppSetting.LogSaveName,
	//	time.Now().Format(setting.AppSetting.TimeFormat),
	//	setting.AppSetting.LogFileExt,
	//)
	return fmt.Sprintf("%s.%s",
		setting.AppSetting.LogSaveName,
		setting.AppSetting.LogFileExt,
	)
}