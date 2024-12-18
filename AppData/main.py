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
