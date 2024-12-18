## PYSTAND_go——python程序发行方法

【参考】：

本项目是在使用Umi-OCR时，对其打包程序的方式PyStand进行探索的过程。

[skywind3000/PyStand: :rocket: Python Standalone Deploy Environment !!](https://github.com/skywind3000/PyStand)

[Umi-OCR_runtime_windows/PyStand_for_UmiOCR at main · hiroi-sora/Umi-OCR_runtime_windows](https://github.com/hiroi-sora/Umi-OCR_runtime_windows/tree/main/PyStand_for_UmiOCR)

### 原理：

在常规开发的Python项目中，venv虚拟环境中的第三方库保存在：**venv\Lib\site-packages** 中。

下载：[Python Releases for Windows | Python.org](https://www.python.org/downloads/windows/)

可以通过使用 **Python 的 Windows embeddable package 版本**，手动配置 **site-packages** 路径，从而在可移植的python中引入 **site-packages** 中的第三方库。

#### 1、start.exe 

- 检测相关文件存在

- 设置环境变量：`"PYTHONDIR": pl.pythonDir` （D:\***\PYSTAND_GO\AppData）

- 运行Python脚本 `AppData\main.py`，重定向输入输出，配置相关参数等

  ```go
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
  ```


#### 2、AppData\main.py 入口函数

- 加载 `AppData\site-packages` 下的第三方库
- 执行真正的 python项目 的入口函数：`run()`

```python
import sys
import os
import copy
import site


def set_env():
    # 保存原始的 Python 搜索路径
    sys.path_origin = [n for n in sys.path]
    python_dir = os.environ['PYTHONDIR']

    for n in ['.', 'lib', 'site-packages', 'runtime']:
        test = os.path.abspath(os.path.join(python_dir, n))
        if os.path.exists(test):
            site.addsitedir(test)
            print(f"添加路径: {test}")


if __name__ == "__main__":
    set_env()
    
    from run import *

    run()
```

------------

### 使用方法：

1. 获取与 待打包python项目 版本一致的 **Python 的 Windows embeddable package 版本**，并解压到：`AppData\runtime` 下。

   下载：[Python Releases for Windows | Python.org](https://www.python.org/downloads/windows/)

2. 找到 待打包python项目 的 **venv虚拟环境** 位置，拷贝其中的 `site-packages` 到 `AppData` 下。

3. 将你的 待打包python项目 的入口文件命名为 `run.py`，入口函数命名为 `run`，拷贝到 `AppData` 下。

4. 在 PYSTAND_GO 根目录下，`go mod init` 初始化项目，`go build main.go` 输出可执行程序 `main.exe`。

5. 分发时，打包 `main.exe`，`AppData`文件夹即可。

```cmd
操作流程：

D:\***\PYSTAND_GO
├─AppData
│  ├─runtime  # (1) 将embeddable版本python解压在这里
│  │  ├─_asyncio.pyd
│  │  ├─...
│  ├─site-packages  # (2) 在常规开发的Python项目中的site-packages拷贝到这
│  │  ├─pip
│  │  ├─numpy
│  │  ├─...
│  ├─main.py  # (4) main.py在embed python环境中导入site-packages第三方库，并运行run.py
│  ├─run.py  # (5) 真正业务代码的入口
│  └─__pycache__
└─start.exe  # (3) start.exe拉起embed python环境，执行main.py
```











