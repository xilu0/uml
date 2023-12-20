# 导入所需库
from graphviz import Digraph

# 创建图表对象
dot = Digraph(comment='The Paxos Algorithm Execution in Dual-Master Node Setup')

# 添加节点和边
dot.node('A', 'Node 1: Proposer/Acceptor')
dot.node('B', 'Node 2: Proposer/Acceptor')

# 描述Paxos算法的执行过程
# 提议阶段
dot.edge('A', 'B', 'Proposal with unique number n')
dot.edge('B', 'A', 'Acknowledgement if n is the highest number seen')

# 接受阶段
dot.edge('A', 'B', 'Requests acceptance of proposal')
dot.edge('B', 'A', 'Accepts proposal if no higher number proposal received')

# 数据写入和冲突解决阶段
dot.edge('A', 'B', 'Write Request 1')
dot.edge('B', 'A', 'Write Request 2')
dot.edge('A', 'B', 'Conflict Detection')
dot.edge('B', 'A', 'Conflict Resolution and Data Synchronization')

# 生成图表
dot.render('paxos_dual_master', format='png', cleanup=True)  # 输出图表文件
dot.view()  # 查看图表文件
