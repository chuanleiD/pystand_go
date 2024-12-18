#!/usr/bin/env python
# -*- coding: UTF-8 -*-
"""
@Project ：AppData 
@File    ：run.py.py
@Author  ：ChenLiang
@Date    ：2024/12/12 下午5:07 
"""


def verify_numpy_installation():
    import numpy as np
    """
    验证 NumPy 是否正确安装并能正常使用的测试脚本
    返回值: 如果所有测试通过返回 True，否则会抛出异常
    """
    try:
        # 步骤 1: 尝试导入 NumPy
        print("步骤 1: 测试 NumPy 导入...")
        print(f"NumPy 版本: {np.__version__}")

        # 步骤 2: 测试基本数组操作
        print("\n步骤 2: 测试基本数组操作...")
        array = np.array([1, 2, 3, 4, 5])
        print("创建数组:", array)
        print("数组求和:", array.sum())

        # 步骤 3: 测试基本数学运算
        print("\n步骤 3: 测试数学运算...")
        print("数组平方:", np.square(array))
        print("数组平均值:", np.mean(array))

        # 步骤 4: 测试多维数组操作
        print("\n步骤 4: 测试多维数组...")
        matrix = np.array([[1, 2], [3, 4]])
        print("2x2矩阵:\n", matrix)
        print("矩阵转置:\n", matrix.T)

        # 步骤 5: 测试随机数生成
        print("\n步骤 5: 测试随机数生成...")
        random_array = np.random.rand(5)
        print("随机数组:", random_array)

        print("\n所有测试通过！NumPy 安装正确且运行正常。")
        return True

    except ImportError:
        print("错误: NumPy 未安装或安装失败。")
        print("请尝试使用以下命令安装 NumPy:")
        print("pip install numpy")
        raise

    except Exception as e:
        print(f"测试过程中发生错误: {str(e)}")
        raise
        
        

def run():
    print("hello, world")
    verify_numpy_installation()


if __name__ == '__main__':
    run()
    
