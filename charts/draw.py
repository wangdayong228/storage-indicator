import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
import seaborn as sns


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

    # 格式化横轴为仅显示时间
    plt.gca().xaxis.set_major_formatter(plt.matplotlib.dates.DateFormatter("%H:%M"))

    # 保存图表
    plt.savefig(f"./out/{chart_name}_line.png")
    plt.show()


# 包含离群值的箱型图
def plot_box_with_outliers(input_csv, chart_name, x_col_index, y_col_index, y_col_log=False):
    # 读取 CSV 文件（假设有 4 列）
    df = pd.read_csv(input_csv, parse_dates=[x_col_index])

    # 提取小时
    df["Hour"] = df[df.columns[x_col_index]].dt.hour

    # 画箱型图，使用最后一列 'Value'
    plt.figure(figsize=(12, 6))
    sns.boxplot(x="Hour", y=df.columns[y_col_index], data=df)

    # 设置标题和标签
    plt.xlabel("Hour")
    plt.ylabel(df.columns[y_col_index])
    plt.title(f"{chart_name}")

    # 显示图表
    plt.grid()
    # 如果需要对数刻度
    if y_col_log:
        plt.yscale("log")

    plt.savefig(f"./out/{chart_name}_boxplot.png")
    plt.show()


# 自定义箱型图（只显示四分位数信息）
# y_col_log 为是否 y 轴指数型表示
def plot_box(input_csv, chart_name, x_col_index, y_col_index, y_col_log=False):
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
        plt.fill_between(
            [row["Hour"] - 0.3, row["Hour"] + 0.3], [row["Q1"], row["Q1"]], [row["Q3"], row["Q3"]], color="skyblue", alpha=0.8
        )

        # 中位线
        plt.plot([row["Hour"] - 0.3, row["Hour"] + 0.3], [row["Median"], row["Median"]], "r", linewidth=2)

        # 胡须（上下限）
        plt.plot([row["Hour"], row["Hour"]], [row["Min"], row["Q1"]], "k--", linewidth=1)
        plt.plot([row["Hour"], row["Hour"]], [row["Q3"], row["Max"]], "k--", linewidth=1)

        # 上下限横线
        plt.plot([row["Hour"] - 0.15, row["Hour"] + 0.15], [row["Min"], row["Min"]], "k-", linewidth=1)
        plt.plot([row["Hour"] - 0.15, row["Hour"] + 0.15], [row["Max"], row["Max"]], "k-", linewidth=1)

    # 设置标题和标签
    plt.xlabel("Hour")
    plt.ylabel(df.columns[y_col_index])
    plt.title(f"{chart_name}")
    plt.xticks(hours)
    plt.grid(True, axis="y", linestyle="--", alpha=0.7)

    # 如果需要对数刻度
    if y_col_log:
        plt.yscale("log")

    # 添加图例
    from matplotlib.patches import Patch

    legend_elements = [
        Patch(facecolor="skyblue", edgecolor="black", alpha=0.8, label="IQR (Q1-Q3)"),
        plt.Line2D([0], [0], color="r", lw=2, label="Median"),
        plt.Line2D([0], [0], color="k", linestyle="--", label="Min/Max Range"),
    ]
    plt.legend(handles=legend_elements, loc="best")

    plt.savefig(f"./out/{chart_name}_boxplot.png")
    plt.show()


def plot_mine_net_rate_minute_average(input_csv, chart_name, timestamp_col_index, value_col_name):
    # 读取 CSV 文件
    df = pd.read_csv(input_csv)

    print("columns", df.columns)
    print("value column", value_col_name)

    # 确保时间戳列转换为 datetime 类型
    df[df.columns[timestamp_col_index]] = pd.to_datetime(df[df.columns[timestamp_col_index]])

    # 过滤数据，只保留前 6 小时的数据
    start_time = df[df.columns[timestamp_col_index]].min()
    end_time = start_time + pd.Timedelta(hours=6)
    df = df[(df[df.columns[timestamp_col_index]] >= start_time) & (df[df.columns[timestamp_col_index]] < end_time)]

    print("value column2", value_col_name)
    # 设置时间戳列为索引
    df.set_index(df.columns[timestamp_col_index], inplace=True)

    # 按分钟重采样并计算平均值
    minute_avg = df[value_col_name].resample("T").mean() * 256 / 1024
    print("value column2", value_col_name)
    # 绘制图表
    plt.figure(figsize=(12, 6))
    plt.plot(minute_avg.index, minute_avg.values, linestyle="-", linewidth=0.5, color="b", label="Average Value per Minute")

    # 设置横轴和纵轴名称
    plt.xlabel("Time (Minutes)")
    print("value column2", value_col_name)
    plt.ylabel(f"{value_col_name}(Avg, Mb/s)")
    plt.title(f"{chart_name}")
    plt.xticks(rotation=45)
    plt.legend()
    plt.grid()

    # 格式化横轴为仅显示时间
    plt.gca().xaxis.set_major_formatter(plt.matplotlib.dates.DateFormatter("%H:%M"))

    # 保存图表
    plt.savefig(f"./out/{chart_name}_minute_avg.png")
    plt.show()


# 全天数据
dir = "./out/hdd"
plot_box(f"{dir}/SyncProgressDiff.csv", "SyncProgressDiff", 0, 3)
plot_box(f"{dir}/MemPoolRefreshRate.csv", "MemPoolRefreshRate", 1, 3)
plot_box_with_outliers(f"{dir}/TxSyncCompleteTimeCost.csv", "TxSyncCompleteTimeCost", 1, 3, True)
plot_box(f"{dir}/SyncTaskBacklog.csv", "SyncTaskBacklog", 0, 1)
plot_box(f"{dir}/MineWork.csv", "MineWork-ScratchPadRate", 0, 2)
plot_mine_net_rate_minute_average(f"{dir}/MineWork.csv", "MineWork-LoadingRate", 0, "LoadingRate")
plot_box(f"{dir}/MineWork.csv", "MineWork-PadMixRate", 0, 6)

# # 少量数据
# plot_box("./out/short/SyncProgressDiff.csv","SyncProgressDiff_short", 0, 3)
# plot_box("./out/short/MemPoolRefreshRate.csv","MemPoolRefreshRate_short", 1, 3)
# plot_box("./out/short/TxSyncCompleteTimeCost.csv","TxSyncCompleteTimeCost_short", 1, 3)
# plot_box("./out/short/SyncTaskBacklog.csv","SyncTaskBacklog_short", 0, 1)
# plot_box("./out/short/Mine
