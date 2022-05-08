package logging


import (
	"GinRouter/pkg/file"
	"fmt"
	"log"
	"os"
	"path"
)


var (
	F *os.File

	LogFileAbsolutePath string

	DefaultPrefix      = ""

)


func init() {
	fmt.Println("初始化日志文件")
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	log.New(F, DefaultPrefix, log.LstdFlags)

	LogFileAbsolutePath = path.Join(filePath, fileName)
	fmt.Printf("日志文件路径为： %s", LogFileAbsolutePath)
}
