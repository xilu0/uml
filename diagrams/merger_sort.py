import matplotlib.pyplot as plt
import numpy as np

# 定义n的范围
n = np.arange(1, 100)

# 计算 O(n) 和 O(n log n)
o_n = n
o_n_log_n = n * np.log(n)

# 创建图形
plt.figure(figsize=(10,6))

# 绘制 O(n) 和 O(n log n) 的图形
plt.plot(n, o_n, label='O(n)', color='blue')
plt.plot(n, o_n_log_n, label='O(n log n)', color='red')

# 添加标签和图例
plt.xlabel('n (Input Size)')
plt.ylabel('Operations')
plt.legend()
plt.title('Growth Comparison between O(n) and O(n log n)')
plt.grid(True)
plt.show()
