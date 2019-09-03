# -*- coding: utf-8 -*
#!/usr/local/bin/python3
import pandas as pd 
import numpy as np

datas = pd.date_range('20190101',periods=6)
df = pd.DataFrame(np.random.randn(24).reshape(6,4),index=datas,columns=['a','b','c','d'])
print(df)

df.to_csv('read.csv')

dd = pd.read_csv('read.csv')
print('dd',dd)