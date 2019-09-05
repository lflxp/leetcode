# -*- coding: utf-8 -*
#!/usr/local/bin/python3
import matplotlib.pyplot as plt
import numpy as np

x = np.linspace(-3,3,50)

y1 = 2*x + 1
y2 = x**2

plt.figure(num=3,figsize=(8,5))
plt.plot(x,y2,label='linear line')
plt.plot(x,y1,color='red',linewidth=1.0,linestyle='--',label='square line')
plt.legend(loc='best') # upper right labels=['up','down'],

# plt.xlim((-1,2))
# plt.ylim((-2,3))
plt.xlabel('X label')
plt.ylabel('Y label')
new_tricks = np.linspace(-1,2,5)
print(new_tricks)
plt.xticks(new_tricks)
plt.yticks([-2, -1.8, -1, 1.22, 3],[r'$really\ bad$', r'$bad$', r'$normal$', r'$good$', r'$really\ good$'])
plt.show()