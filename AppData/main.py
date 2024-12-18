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
    
    import os
    
    # 获取当前脚本的绝对路径
    current_file_path = os.path.abspath(__file__)

    # 获取当前脚本所在的目录
    current_dir = os.path.dirname(current_file_path)

    # 将当前工作目录切换到脚本所在目录
    os.chdir(current_dir)

    # 验证当前工作目录
    print(f"当前工作目录: {os.getcwd()}")

    from run import *

    run()
