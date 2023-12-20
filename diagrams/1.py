import matplotlib.pyplot as plt
import networkx as nx
import matplotlib
matplotlib.rcParams['font.sans-serif'] = ['SimHei']  # 指定默认字体为新宋体。
matplotlib.rcParams['axes.unicode_minus'] = False  # 解决保存图像是负号 '-' 显示为方块的问题。

# Create a directed graph to visually represent the article's main ideas
G2 = nx.DiGraph()

# Define nodes and edges based on the article's structure
nodes2 = {
    "CurrentStatus": "当前状况：Go 开发工程师",
    "HighGoal": "问题：架构师目标是否过高？",
    "RealisticGoal": "更合情理的目标：高级 Go 开发工程师",
    "Advantages": "优势",
    "AchieveGoal": "如何达成目标",
    "Conclusion": "总结"
}

edges2 = [
    ("CurrentStatus", "HighGoal", "思考长远发展"),
    ("HighGoal", "RealisticGoal", "权衡可行性"),
    ("RealisticGoal", "Advantages", "明确优势"),
    ("RealisticGoal", "AchieveGoal", "规划路径"),
    ("AchieveGoal", "Conclusion", "实施方案"),
    ("Advantages", "Conclusion", "综合考量")
]

# Add nodes and edges to the graph
G2.add_nodes_from(nodes2.keys())
G2.add_edges_from((src, dest, {'label': label}) for src, dest, label in edges2)

# Draw the graph
pos2 = nx.spring_layout(G2, seed=43)  # positions for all nodes
labels2 = nx.get_edge_attributes(G2, 'label')

plt.figure(figsize=(14, 10))
nx.draw(G2, pos2, with_labels=True, node_color='lightgreen', font_size=12, font_color='black', node_size=3000)
nx.draw_networkx_edge_labels(G2, pos2, edge_labels=labels2, font_size=10)

plt.title("从 Go 开发工程师到高级 Go 开发工程师：目标设定与职业规划", fontsize=16)
plt.show()
