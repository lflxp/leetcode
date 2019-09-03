# -*- coding: utf-8 -*
#!/usr/bin/python3
import pandas as pd 
import numpy as np

# 基本介绍
# s = pd.Series([1,3,6,np.nan,44,1])
# # print(s)

# dates = pd.date_range('20190101',periods=6)
# df = pd.DataFrame(np.random.randn(6,4),index=dates,columns=['a','b','c','d'])
# print(df)

# df1 = pd.DataFrame(np.random.rand(6,4))
# print(df1)
# print(df['b'])

# df2 = pd.DataFrame(np.arange(12).reshape(3,4))
# print(df2)

# df3 = pd.DataFrame({'A' : 1.,
#                     'B' : pd.Timestamp('20130102'),
#                     'C' : pd.Series(1,index=list(range(4)),dtype='float32'),
#                     'D' : np.array([3] * 4,dtype='int32'),
#                     'E' : pd.Categorical(["test","train","test","train"]),
#                     'F' : 'foo'})
# print(df3)                    
# print(df3.dtypes,df3.index,df3.columns)
# print(df3.values)
# print(df3.T)
# print(df3.sort_index(axis=1,ascending=False))
# print(df3.sort_index(axis=0,ascending=False))
# print(df3.mean(axis=1))
# print(df3.mean(axis=0))

# print(df3.sort_values(by='E'))

# pandas 选择数据
dates = pd.date_range('20190101',periods=6)
# print(dates)
sd = pd.DataFrame(np.arange(24).reshape((6,4)),index=dates,columns=['A','B','C','D'])
# print(sd)
# print(sd['A'],sd.A)
# print(sd[0:3])
# print(sd['20190101':'20190103'])

# select by label: loc
# print(sd.loc['20190102'])
# print(sd.loc[:,['A','B']])
# print(sd.loc['20190103',['A','B']])

# select by position: iloc
# print(sd.iloc[[1,3,5],1:3])

# mixed selection: ix 不推荐
# print(sd.ix[:3,['A','C']])

# Boolean indexing
# print(sd[sd.A > 17])

sd.iloc[2,2] = 212
print(sd)

sd.loc['20190104','B'] = 998
print(sd)

sd.loc['20190104',['A','B']] = [1,2]
print(sd)

sd[sd.B > 4] = 0
print(sd)

sd.A[sd.A < 1] = 2
print(sd)

sd['F'] = np.nan
print(sd)

sd.F = 9
print(sd)

sd['E'] = pd.Series([1,2,3,4,5,6],index=dates)
print(sd)