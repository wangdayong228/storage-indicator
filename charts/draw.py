import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

# 折线图
def plot_line(input_csv, chart_name, x_col_index, y_col_index):
    # 读取 CSV 文件
    df = pd.read_csv(input_csv)

    # 确保时间戳列转换为 datetime 类型（如果需要）
    df.iloc[:, x_col_index] = pd.to_datetime(df.iloc[:, x_col_index])

    # 绘制图表
    plt.figure(figsize=(12, 6))
    plt.plot(
        df.iloc[:, x_col_index], df.iloc[:, y_col_index], linestyle="-", linewidth=0.5, color="b", label=df.columns[y_col_index]
    )

    # 设置横轴和纵轴名称
    plt.xlabel(df.columns[x_col_index])
    plt.ylabel(df.columns[y_col_index])
    plt.title(f"{chart_name}")
    plt.xticks(rotation=45)
    plt.legend()
    plt.grid()

    # 保存图表
    plt.savefig(f"./out/{chart_name}.png")
    plt.show()

# 自定义箱型图（只显示四分位数信息）
def plot_box(input_csv, chart_name, x_col_index, y_col_index):
    # 读取 CSV 文件
    df = pd.read_csv(input_csv, parse_dates=[x_col_index])

    # 提取小时
    df["Hour"] = df[df.columns[x_col_index]].dt.hour
    
    # 准备图表
    plt.figure(figsize=(12, 6))
    
    # 按小时分组并计算统计量
    hours = sorted(df["Hour"].unique())
    stats_data = []
    
    for hour in hours:
        hour_data = df[df["Hour"] == hour][df.columns[y_col_index]]
        if not hour_data.empty:
            # 计算简单的四分位数统计
            min_val = hour_data.min()
            q1 = hour_data.quantile(0.25)
            median = hour_data.median()
            q3 = hour_data.quantile(0.75)
            max_val = hour_data.max()
            stats_data.append([hour, min_val, q1, median, q3, max_val])
    
    stats_df = pd.DataFrame(stats_data, columns=["Hour", "Min", "Q1", "Median", "Q3", "Max"])
    
    # 绘制自定义箱型图
    for i, row in stats_df.iterrows():
        # 箱体（Q1-Q3）
        plt.fill_between([row["Hour"]-0.3, row["Hour"]+0.3], [row["Q1"], row["Q1"]], 
                         [row["Q3"], row["Q3"]], color='skyblue', alpha=0.8)
        
        # 中位线
        plt.plot([row["Hour"]-0.3, row["Hour"]+0.3], [row["Median"], row["Median"]], 'r', linewidth=2)
        
        # 胡须（上下限）
        plt.plot([row["Hour"], row["Hour"]], [row["Min"], row["Q1"]], 'k--', linewidth=1)
        plt.plot([row["Hour"], row["Hour"]], [row["Q3"], row["Max"]], 'k--', linewidth=1)
        
        # 上下限横线
        plt.plot([row["Hour"]-0.15, row["Hour"]+0.15], [row["Min"], row["Min"]], 'k-', linewidth=1)
        plt.plot([row["Hour"]-0.15, row["Hour"]+0.15], [row["Max"], row["Max"]], 'k-', linewidth=1)
    
    # 设置标题和标签
    plt.xlabel("Hour")
    plt.ylabel(df.columns[y_col_index])
    plt.title(f"{chart_name}")
    plt.xticks(hours)
    plt.grid(True, axis='y', linestyle='--', alpha=0.7)
    
    # 添加图例
    from matplotlib.patches import Patch
    legend_elements = [
        Patch(facecolor='skyblue', edgecolor='black', alpha=0.8, label='IQR (Q1-Q3)'),
        plt.Line2D([0], [0], color='r', lw=2, label='Median'),
        plt.Line2D([0], [0], color='k', linestyle='--', label='Min/Max Range')
    ]
    plt.legend(handles=legend_elements, loc='best')
    
    plt.savefig(f"./out/{chart_name}_boxplot.png")
    plt.show()

# 全天数据
plot_box("./out/SyncProgressDiff.csv", "SyncProgressDiff", 0, 3)
plot_box("./out/MemPoolRefreshRate.csv", "MemPoolRefreshRate", 1, 3)
plot_box("./out/TxSyncCompleteTimeCost.csv", "TxSyncCompleteTimeCost", 1, 3)
plot_box("./out/SyncTaskBacklog.csv", "SyncTaskBacklog", 0, 1)
plot_box("./out/MineWork.csv", "MineWork", 0, 1)

# # 少量数据
# plot_box("./out/short/SyncProgressDiff.csv","SyncProgressDiff_short", 0, 3)
# plot_box("./out/short/MemPoolRefreshRate.csv","MemPoolRefreshRate_short", 1, 3)
# plot_box("./out/short/TxSyncCompleteTimeCost.csv","TxSyncCompleteTimeCost_short", 1, 3)
# plot_box("./out/short/SyncTaskBacklog.csv","SyncTaskBacklog_short", 0, 1)
# plot_box("./out/short/Mine