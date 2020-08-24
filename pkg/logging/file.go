package logging

import (
	"fmt"
	"go-gin-example/pkg/file"
	"go-gin-example/pkg/setting"
	"os"
	"time"
)

//var (
//	LogSavePath = "runtime/logs/"
//	LogSaveName = "log"
//	LogFileExt = "log"
//	TimeFormat = "20060102"
//)




func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt)

	return fmt.Sprintf("%s%s",prefixPath, suffixPath)
}


func openLogFile(fileName, filePath string) (*os.File, error) {
	dir,err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.checkPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExitMkDir src: %s, err: %v", src, err)
	}

	f,err := file.Open(src + fileName, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile: %v", err)
	}
	return f, nil
}
/*
func openLogFile(filePath string) *os.File {
	// os.Stat：返回文件信息结构描述文件。如果出现错误，会返回*PathError
	_, err := os.Stat(filePath)
	switch  {
		// os.IsNotExist：能够接受ErrNotExist、syscall的一些错误，它会返回一个布尔值，能够得知文件不存在或目录不存在
	case os.IsNotExist(err):
		// os.IsPermission：能够接受ErrPermission、syscall的一些错误，它会返回一个布尔值，能够得知权限是否满足
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}
	// os.OpenFile：调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于 I/O。如果出现错误，则为*PathError
	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile:%v", err)
	}
	return handle
}
 */

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}