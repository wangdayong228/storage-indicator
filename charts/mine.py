import matplotlib.pyplot as plt
import pandas as pd

# 读取 CSV 文件
df = pd.read_csv('out/MineWork.csv')

# 将 Timestamp 列转换为 datetime 类型
df['Timestamp'] = pd.to_datetime(df['Timestamp'])

# 创建图形
fig, ax1 = plt.subplots(figsize=(12, 6))

# 绘制第一个 y 轴的数据
ax1.plot(df['Timestamp'], df['ScratchPad'], 'g-', label='ScratchPad')
ax1.set_xlabel('Timestamp')
ax1.set_ylabel('ScratchPad', color='g')
ax1.tick_params(axis='y', labelcolor='g')

# 创建第二个 y 轴
ax2 = ax1.twinx()
ax2.plot(df['Timestamp'], df['Loading'], 'r-', label='Loading')
ax2.set_ylabel('Loading', color='r')
ax2.tick_params(axis='y', labelcolor='r')

# 创建第三个 y 轴
ax3 = ax1.twinx()
ax3.spines['right'].set_position(('outward', 60))  # 将第三个 y 轴向右移动
ax3.plot(df['Timestamp'], df['PadMix'], 'b-', label='PadMix')
ax3.set_ylabel('PadMix', color='b')
ax3.tick_params(axis='y', labelcolor='b')

# 创建第四个 y 轴
ax4 = ax1.twinx()
ax4.spines['right'].set_position(('outward', 120))  # 将第四个 y 轴向右移动
ax4.plot(df['Timestamp'], df['Hit'], 'y-', label='Hit')
ax4.set_ylabel('Hit', color='y')
ax4.tick_params(axis='y', labelcolor='y')

# 添加图例
fig.legend(loc='upper left', bbox_to_anchor=(0.1, 0.9))

# 设置图表标题
plt.title('Multi Y-Axis Plot for CSV Data')

# 显示图形
plt.show()