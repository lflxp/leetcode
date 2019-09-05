# -*- coding: utf-8 -*
#!/usr/local/bin/python3
import matplotlib.pyplot as plt
import numpy as np

x = np.linspace(-1,1,50)
print(x)
# y = 2*x + 1
y = x**2
print(y)
plt.plot(x,y)
plt.show()