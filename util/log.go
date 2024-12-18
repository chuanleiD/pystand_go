package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Init(AppData string) error {
	// 创建构建目录
	if _, err := os.Stat("build"); os.IsNotExist(err) {
		err = os.Mkdir("build", 0755)
		if err != nil {
			return fmt.Errorf("构建目录%s创建失败: %v", "build", err)
		}
	}

	dirPath := filepath.Join("build", AppData)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("构建目录%s创建失败: %v", dirPath, err)
		}
	}

	return nil
}

// SetupLogger 初始化日志设置
func SetupLogger() (*os.File, error) {
	// 获取当前位置
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("获取可执行文件路径失败: %v", err)
	}
	homeDir := filepath.Dir(exePath)

	// 创建日志文件名称
	timestamp := time.Now().Format("20060102")
	err = os.MkdirAll("log", 0755)
	if err != nil {
		return nil, fmt.Errorf("日志目录创建失败: %v", err)
	}
	logFilePath := filepath.Join(homeDir, "log", "log_"+timestamp+".txt")

	// 打开日志文件
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开日志文件失败: %v", err)
	}

	// 创建一个多重写入器
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 设置日志输出到多重写入器
	log.SetOutput(multiWriter)

	// 设置日志前缀和标志
	log.SetPrefix("【INFO】")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return logFile, nil
}
