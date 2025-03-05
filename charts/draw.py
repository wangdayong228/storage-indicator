# import pandas as pd
# import matplotlib.pyplot as plt

# # 读取 CSV 文件
# csv_file = "output.csv"
# df = pd.read_csv(csv_file)

# # 确保时间戳列转换为 datetime 类型（如果需要）
# df.iloc[:, 0] = pd.to_datetime(df.iloc[:, 0])

# # 绘制图表
# plt.figure(figsize=(12, 6))
# plt.plot(df.iloc[:, 0], df.iloc[:, 3], marker='o', linestyle='-', color='b', label="Diff")

# # 设置横轴格式
# plt.xlabel("Timestamp")
# plt.ylabel("Diff")
# plt.title("Timestamp vs Diff")
# plt.xticks(rotation=45)
# plt.legend()
# plt.grid()

# # 保存图表
# plt.savefig("sync_diff.png")
# plt.show()

import pandas as pd
import matplotlib.pyplot as plt

# 定义函数以支持参数输入
def plot(input_csv, chart_name, x_col_index, y_col_index):
    # 读取 CSV 文件
    df = pd.read_csv(input_csv)

    # 确保时间戳列转换为 datetime 类型（如果需要）
    df.iloc[:, x_col_index] = pd.to_datetime(df.iloc[:, x_col_index])

    # 绘制图表
    plt.figure(figsize=(12, 6))
    plt.plot(df.iloc[:, x_col_index], df.iloc[:, y_col_index], marker='o', linestyle='-', color='b', label=df.columns[y_col_index])

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

# 全天数据
# plot("./out/SyncProgressDiff.csv","SyncProgressDiff", 0, 3)
# plot("./out/MemPoolRefreshRate.csv","MemPoolRefreshRate", 0, 3)
# plot("./out/TxSyncCompleteTimeCost.csv","TxSyncCompleteTimeCost", 1, 3)
# plot("./out/SyncTaskBacklog.csv","SyncTaskBacklog", 0, 1)
plot("./out/MineWork.csv","MineWork", 0, 1)

# 少量数据
# plot("./out/short/SyncProgressDiff.csv","SyncProgressDiff_short", 0, 3)
# plot("./out/short/MemPoolRefreshRate.csv","MemPoolRefreshRate_short", 1, 3)
# plot("./out/short/TxSyncCompleteTimeCost.csv","TxSyncCompleteTimeCost_short", 1, 3)
# plot("./out/short/SyncTaskBacklog.csv","SyncTaskBacklog_short", 0, 1)
# plot("./out/short/MineWork.csv","MineWork_short", 0, 1)
