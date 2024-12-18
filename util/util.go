package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

// PythonLauncher 结构体存储运行时信息
type PythonLauncher struct {
	exePath          string // 当前可执行文件路径
	homeDir          string // 程序所在目录
	pythonDir        string // Python运行时目录
	runtimePath      string // Python运行时路径
	sitePackagesPath string // Python第三方库路径
	scriptPath       string // 主脚本路径
	pythonDLL        string // python3X.dll 路径
}

// ShowDetail 显示详细信息
func (pl *PythonLauncher) ShowDetail() {
	log.Println("exePath:", pl.exePath)
	log.Println("homeDir:", pl.homeDir)
	log.Println("pythonDir:", pl.pythonDir)
	log.Println("runtimePath:", pl.runtimePath)
	log.Println("sitePackagesPath:", pl.sitePackagesPath)
	log.Println("scriptPath:", pl.scriptPath)
	log.Println("pythonDLL:", pl.pythonDLL)
	return
}

// NewPythonLauncher 初始化启动器，获取Python运行时的各种路径信息
func NewPythonLauncher(AppData string) (*PythonLauncher, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("获取可执行文件路径失败: %v", err)
	}

	homeDir := filepath.Dir(exePath)

	pl := &PythonLauncher{
		exePath:          exePath, // 程序所在目录/exePath
		homeDir:          homeDir, // 程序所在目录=homeDir
		pythonDir:        filepath.Join(homeDir, AppData),
		runtimePath:      filepath.Join(homeDir, AppData, "runtime"),
		sitePackagesPath: filepath.Join(homeDir, AppData, "site-packages"),
		scriptPath:       filepath.Join(homeDir, AppData, "main.py"),
		pythonDLL:        filepath.Join(homeDir, AppData, "runtime", "python3.dll"),
	}

	// 展示Python运行时的各种路径信息
	pl.ShowDetail()

	return pl, nil
}

// CheckEnvironment 检查环境
func (pl *PythonLauncher) CheckEnvironment() error {

	if _, err := os.Stat(pl.exePath); os.IsNotExist(err) {
		return fmt.Errorf("exePath: %s", pl.exePath)
	}
	if _, err := os.Stat(pl.homeDir); os.IsNotExist(err) {
		return fmt.Errorf("homeDir: %s", pl.homeDir)
	}
	if _, err := os.Stat(pl.pythonDir); os.IsNotExist(err) {
		return fmt.Errorf("pythonDir: %s", pl.pythonDir)
	}
	if _, err := os.Stat(pl.runtimePath); os.IsNotExist(err) {
		return fmt.Errorf("runtimePath: %s", pl.runtimePath)
	}
	if _, err := os.Stat(pl.sitePackagesPath); os.IsNotExist(err) {
		return fmt.Errorf("sitePackagesPath: %s", pl.sitePackagesPath)
	}
	if _, err := os.Stat(pl.scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("scriptPath: %s", pl.scriptPath)
	}
	if _, err := os.Stat(pl.pythonDLL); os.IsNotExist(err) {
		return fmt.Errorf("pythonDLL: %s", pl.pythonDLL)
	}

	return nil
}

// SetupEnvironment 设置环境变量
func (pl *PythonLauncher) SetupEnvironment() error {
	env := map[string]string{
		"PYTHONDIR": pl.pythonDir,
	}
	for k, v := range env {
		fmt.Println(k, v)
	}

	for k, v := range env {
		if err := os.Setenv(k, v); err != nil {
			return fmt.Errorf("设置环境变量失败 %s: %v", k, err)
		}
	}

	return nil
}

// RunScript 运行Python脚本
func (pl *PythonLauncher) RunScript() error {
	cmd := exec.Command(filepath.Join(pl.runtimePath, "python.exe"), "-I", "-s", "-S", pl.scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 如果是Windows GUI程序，需要特殊处理
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}
	}

	return cmd.Run()
}
