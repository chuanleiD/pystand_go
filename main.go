package main

import (
	"log"
	"os"
	"pystand_go/util"
)

const AppData = "AppData"

func main() {
	// 初始化
	err := util.Init(AppData)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 设置日志
	logFile, err := util.SetupLogger()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer logFile.Close()

	// 创建启动器实例
	launcher, err := util.NewPythonLauncher(AppData)
	if err != nil {
		log.Printf("初始化失败: %v", err)
		os.Exit(1)
	}

	// 检查环境
	if err := launcher.CheckEnvironment(); err != nil {
		log.Printf("环境检查失败: %v", err)
		os.Exit(1)
	}

	// 设置环境变量
	if err := launcher.SetupEnvironment(); err != nil {
		log.Printf("环境设置失败: %v", err)
		os.Exit(1)
	}

	// 运行脚本
	if err := launcher.RunScript(); err != nil {
		log.Printf("脚本运行失败: %v", err)
		os.Exit(1)
	}
}
