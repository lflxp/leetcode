#!/usr/bin/env python3
import numpy as np

def test1():
    print(np.eye(4))
    a = np.array([1,2,3])
    print(a)
    # 多于一个维度
    b= np.array([[1,2,3],[3,4,5]]) 
    print(b)
    # 最小维度  
    c = np.array([1,2,3,4,5],ndmin=3) 
    print(c)
    # dtype 参数  
    d = np.array([1,2,3],dtype=complex,ndmin=2)
    print(d)

# test1()

# 数据类型 https://www.runoob.com/numpy/numpy-dtype.html
def typetest():
    student = np.dtype([('name','S20'),('age','i1'),('marks','f4')]) # S字符串 i 整型 f 浮点型
    a = np.array([('abc',21,50),('xyz',18,75)],dtype=student)
    print(a)
    print(a.ndim)
# typetest()

# 数组属性
def dataOption():
    a = np.arange(24)
    print(a,a.ndim)
    # 调整大小
    b = a.reshape(2,4,3)
    print(b,b.ndim)
    # ndarray.shape 表示数组的维度，返回一个元组，这个元组的长度就是维度的数目，即 ndim 属性(秩)。比如，一个二维数组，其维度表示"行数"和"列数"。
    print(b.shape)

# dataOption()

# 调整数组大小
def shapeSize():
    a = np.array([[1,2,3],[4,5,6]])
    a.shape = (3,2)
    print(a)
    b = a.reshape(2,3)
    print(b)
    # ndarray.itemsize 以字节的形式返回数组中每一个元素的大小
    print(a.itemsize,b.itemsize)
    # ndarray.flags 返回 ndarray 对象的内存信息
    print(b.flags)
# shapeSize()

# 创建指定大小的数组，数组元素以 0 来填充：
def zeros():
    # 默认为浮点数
    x = np.zeros(5)
    print(x)
    # 设置类型为整数
    y = np.zeros((5,), dtype=np.int)
    print(y)
    # 自定义类型
    z = np.zeros((2,2), dtype = [('x', 'i4'), ('y', 'i4')])  
    print(z)
# zeros()

# 创建指定形状的数组，数组元素以 1 来填充：
# https://www.runoob.com/numpy/numpy-array-creation.html
def ones():
    # 默认为浮点数
    x = np.ones(5) 
    print(x)
    
    # 自定义类型
    x = np.ones([2,2], dtype = int)
    print(x)
ones()