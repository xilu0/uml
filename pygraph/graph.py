from diagrams import Diagram
from diagrams.generic.blank import Blank

with Diagram("系统架构模式与软件架构模式比较", show=True):
    with Diagram("系统架构模式"):
        sys_a = Blank("组织整体系统结构")
        sys_b = Blank("关注硬件、网络等")
        sys_c = Blank("包括多个子系统的协作")
        sys_d = Blank("通常涉及底层资源分配")
        sys_a - sys_b - sys_c - sys_d

    with Diagram("软件架构模式"):
        soft_a = Blank("组织软件模块和组件")
        soft_b = Blank("关注代码结构和互动")
        soft_c = Blank("强调可扩展、可维护性")
        soft_d = Blank("常用于单一应用或服务")
        soft_a - soft_b - soft_c - soft_d
