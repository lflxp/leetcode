# -*- coding: utf-8 -*
#!/usr/bin/env python3
import numpy as np 
from matplotlib import pyplot as plt 
 
# x = np.arange(1,11) 
x = np.linspace(-1,1,200)
np.random.shuffle(x)
y =  0.2  * x +  5 + np.random.normal(0,0.05,(200,))
plt.title("Matplotlib demo") 
plt.xlabel("x axis caption") 
plt.ylabel("y axis caption") 
# plt.plot(x,y) # 连线图 
plt.scatter(x,y) # 散点图
plt.show()
