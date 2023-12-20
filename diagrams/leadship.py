import matplotlib.pyplot as plt
import matplotlib

# 指定支持中文的字体
matplotlib.rcParams['font.sans-serif'] = ['SimHei']
matplotlib.rcParams['axes.unicode_minus'] = False  # 解决负号'-'显示为方块的问题


# 定义层次和颜色
levels = [
    "1. 生存领导\nSurvival Leadership",
    "2. 关系领导\nRelationship Leadership",
    "3. 自我领导\nSelf Leadership",
    "4. 管理领导\nManagerial Leadership",
    "5. 动态领导\nMotivational Leadership",
    "6. 战略领导\nStrategic Leadership",
    "7. 全球领导\nGlobal Leadership",
]
colors = ['#FFDDC7', '#FFCC9C', '#FFAA6E', '#FF7733', '#FF4D00', '#CC3D00', '#992D00']

# 创建图形
fig, ax = plt.subplots()

# 绘制每一层
for i, (level, color) in enumerate(zip(levels, colors)):
    ax.fill_between([0, 7-i*2, 7-(i+1)*2, 0], [7-i, 7-(i+1), 7-(i+1), 7-i], color=color)
    ax.text(3.5, 7-i-0.5, level, ha='center', va='center', fontsize=10)

# 移除坐标轴
ax.axis('off')

# 显示图形
plt.show()
