#!/usr/local/bin/python3
import numpy as np

# array = np.array([[1,2,3],[4,5,6]])
# print(array)
# print(array.ndim)
# print(array.shape)
# print(array.size)

# a1 = np.array([1,3,4],dtype=np.float)
# print(a1,a1.dtype)

# a2 = np.empty((3,4),dtype=np.float)
# print(a2,a2.ndim)

# a3 = np.arange(10,20,2) # 10-19的数据，步长为2
# print(a3)

# a4 = np.zeros((5,5))
# print(a4)

# a5 = np.ones((5,5))
# print(a5)

# a6 = np.arange(20).reshape((4,5))
# print(a6)

# a7 = np.linspace(1,1000,500).reshape((25,20))
# print(a7)

# b1 = np.array([1,3,4,5])
# b2 = np.arange(4)
# print(b1-b2)
# print(b2**2)
# print(b2*np.sin(b1))
# print(b1<3)
# print(b1==3,b2==3)
# b3 = np.arange(5,9).reshape((2,2))
# print(b3)
# b4 = np.array([[1,1],[0,1]])

# c1 = b3*b4
# print(c1)
# # 5 6   1  1
# # 7 8   0  1  

# # 5*1+6*0=5
# # 5*1+6*1=11
# # 7*1+8*0=7
# # 7*1+8*1=15
# c2 = np.dot(b3,b4)
# print(c2)
# c3 = b3.dot(b4)
# print(c3)



# d1 = np.random.random((2,4))
# print(d1)
# print(np.sum(d1,axis=1)) # 求值于每一行
# print(np.max(d1,axis=0)) # 求值于每一列
# print(np.min(d1))

e1 = np.arange(23,0,-2).reshape((3,4))
print(e1)
# print(np.argmax(e1))
# print(np.argmin(e1))
# print(np.mean(e1))
# print(np.average(e1))
# print(np.median(e1))
# print(np.cumsum(e1))
# print(np.diff(e1))
# print(np.nonzero(e1))
# print(np.sort(e1))
# print(np.transpose(e1))
# print(e1.T)
# print(np.clip(e1,3,19))

print(e1[1][1])
print(e1[1,1])
print(e1.T[0,0])

for column in e1.T:
    print(column)

print(e1[1,1:3])
print(e1[2,:])
print(e1.flatten())
print(e1.T.flatten())

for item in e1.flat: # flat是迭代器
    print(item)