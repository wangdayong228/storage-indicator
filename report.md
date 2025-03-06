# 数据存储系统性能测试报告

## 一、概述

本次测试旨在全面评估基于区块链技术的分布式存储系统在长期持续高负载条件下的稳定性、可靠性及性能表现。为最大程度地模拟真实生产环境中的系统运行负载，本次测试持续进行 100 天，总计约 2400 小时的连续不间断操作。通过密集的读写请求模拟真实业务场景，重点考察系统长期运行中的稳定性、吞吐能力、响应延迟、资源利用效率以及数据一致性等关键性能指标。测试过程中，通过严格的监控手段记录系统在不同负载阶段的表现，旨在为系统上线运行提供科学、客观的性能参考依据和优化建议。

## 二、测试环境

为确保测试的准确性和可复现性，测试环境采用标准化、规模化的云计算平台搭建。具体配置如下：

### 1. 硬件环境

- **服务器**：阿里云虚拟服务器
- **处理器（CPU）**：Intel(R) Xeon(R) Platinum 8575C 虚拟 CPU，8 核 vCPU
- **内存**：32 GB DDR4 RAM
- **存储设备**：企业级 SSD 存储，总容量 1.2 TB
- **网络带宽**：100 Mbps 上下行带宽

### 2. 软件环境

- **操作系统**：Ubuntu 22.04 LTS，内核版本 Linux 6.8.0-40-generic，已进行安全更新和性能优化
- **被测系统**：基于区块链技术构建的分布式存储系统，由送测单位提供并部署在我方测试环境中
- **测试工具**：
  - 基于区块链的分布式存储系统专用性能评估工具，用于生成真实的读写操作负载
  - 阿里云系统运行监控平台，用于实时监控节点资源使用情况，包括 CPU 利用率、内存占用、存储空间使用情况、网络流量及系统运行状态等指标


<!-- ### 3. 测试数据

- **数据类型**: [数据类型，如文本、图片、视频等] -->

## 三、测试方法

为准确评估系统在真实业务场景下的表现，测试方法设计充分考虑了基于区块链分布式存储的业务特点，具体测试方案如下：

### 1. 测试场景描述

本次测试对象为送测单位提供的基于区块链技术的分布式存储系统，系统由区块链网络和分布式存储网络两部分组成。送测单位提供完整的区块链测试环境和分布式存储网络测试环境，我方在此基础上运行并维护一个独立的分布式存储节点，加入送测单位所提供的测试网络。

测试期间，通过模拟真实场景需求，节点持续向网络提交数据存储任务（写入请求），同时响应来自区块链网络下发的存储挖矿任务（读取请求），以模拟实际运营环境的复杂性和高并发特征。具体负载设计如下：

- **数据写入负载**：
  - 每间隔 10 分钟自动向网络提交一次数据存储请求；
  - 每次请求的数据大小、内容类型与实际业务场景保持一致，确保负载真实性；

- **数据读取负载（区块链存储挖矿任务）**：
  - 区块链网络每隔 20 分钟自动向节点下发一次存储挖矿任务；
  - 节点以最大能力加载并验证所存储的数据，以模拟真实存储挖矿业务所需的读取及验证操作；

### 2. 测试过程和监控方式

为确保测试结果的可靠性和有效性，测试过程中实施了严格的监控和记录机制：

- **运行监控**：利用阿里云监控平台实时采集系统运行的关键性能指标，包括但不限于 CPU 使用率、内存占用率、磁盘存储使用情况、网络带宽使用情况、系统吞吐量以及响应延迟等信息；

- **数据记录**：测试期间定期导出并保存系统运行状态数据、日志文件及性能指标数据，确保后续分析的准确性和完整性；

<!-- - **异常情况处理**：测试过程中如遇到系统异常情况，及时记录异常发生时刻、现象、持续时长及后续恢复情况，为后续分析问题根因及优化提供客观依据。 -->


## 四、测试结果

### 1. 同步进度差
![`SyncProgressDiff.png`](./out/SyncProgressDiff_boxplot.png) 

同步进度差是衡量基于区块链的分布式存储系统节点与区块链主链同步进度差异的重要技术指标。它反映了节点能否及时有效地跟踪区块链网络上的最新状态，并实时处理链上订单和交易任务的能力。同步进度差异过大可能导致节点无法及时获取链上数据，进而影响系统的一致性、稳定性和业务连续性。

本次持续性测试期间，采用自动化监控工具对节点同步进度差异进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 124718 次同步进度差异指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，无论节点处于持续读写负载还是空闲状态下，每小时同步进度差异的 75% 分位数均稳定在 2 个区块以内，体现出节点在大部分时间内都能高效地追踪主链最新状态。此外，在整个测试过程中，节点所记录的同步进度差异最大值也未超过 4 个区块，这表明即使在高负载或网络波动的情况下，节点仍能较快跟进至最新区块高度。

上图选取了测试期间具有代表性的一天数据，绘制出节点同步进度差异随时间变化的趋势图。从该图可以清晰观察到同步进度差异在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明节点与区块链主链之间的同步进度差异始终保持在较低水平，节点能够可靠、高效地追踪链上状态，满足实际业务场景中对于及时性和稳定性的需求。

### 2. 内存池刷新效率
![`MemPoolRefreshRate.png`](./out/MemPoolRefreshRate_boxplot.png) 

内存池刷新效率是衡量系统在内存管理方面性能的重要指标。反映了系统在处理内存数据时的响应速度和效率。内存池刷新效率低下可能导致系统在处理高并发任务时出现延迟，进而影响整体性能。

本次持续性测试期间，采用自动化监控工具对内存池刷新时间进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 51844 次内存池刷新时间指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，无论节点处于持续读写负载还是空闲状态下，内存池刷新时间的 75% 分位数均稳定在 10 微秒以内，体现出系统在大部分时间内都能高效地处理内存数据。此外，在整个测试过程中，内存池刷新时间的最大值也未超过 60 微秒，这表明即使在高负载或网络波动的情况下，系统仍能保持较高的内存管理效率。

上图选取了测试期间具有代表性的一天数据，绘制出内存池刷新时间随时间变化的趋势图。从该图可以清晰观察到内存池刷新时间在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在内存池刷新效率方面表现优异，能够高效处理内存数据，满足实际业务场景中对于响应速度和稳定性的需求。

### 3. 事务同步时间

![`TxSyncCompleteTimeCost.png`](./out/TxSyncCompleteTimeCost_boxplot.png) 

事务同步时间是评估系统在处理链上事务时的性能指标。反映了系统在处理链上事务时的响应速度和效率。事务同步时间过长可能导致系统在处理高并发任务时出现延迟，进而影响整体性能。

本次持续性测试期间，采用自动化监控工具对事务同步时间进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 4910 次事务同步时间指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，无论节点处于持续读写负载还是空闲状态下，事务同步时间的 75% 分位数均稳定在 10 秒以内，体现出系统在大部分时间内都能高效地处理链上事务。<!-- 此外，在整个测试过程中，事务同步时间的最大值也未超过 20 秒，这表明即使在高负载或网络波动的情况下，系统仍能保持较高的事务处理效率。-->


上图选取了测试期间具有代表性的一天数据，绘制出事务同步时间随时间变化的趋势图。从该图可以清晰观察到事务同步时间在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在事务同步时间方面表现优异，能够高效处理链上事务，满足实际业务场景中对于响应速度和稳定性的需求。


### 4. 同步任务积压

![`SyncTaskBacklog.png`](./out/SyncTaskBacklog_boxplot.png) 

同步任务积压是衡量系统在任务调度方面性能的重要指标。反映了系统在处理任务时的响应速度和效率。同步任务积压过多可能导致系统在处理高并发任务时出现延迟，进而影响整体性能。

本次持续性测试期间，采用自动化监控工具对同步任务积压进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 28801 次同步任务积压指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，无论节点处于持续读写负载还是空闲状态下，同步任务积压的数量 100% 分位数均稳定在 16 个以下，体现出系统在全部时间内都能高效地处理任务。

上图选取了测试期间具有代表性的一天数据，绘制出同步任务积压随时间变化的趋势图。从该图可以清晰观察到同步任务积压在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在同步任务积压方面表现优异，能够高效处理任务，满足实际业务场景中对于响应速度和稳定性的需求。

### 5. 挖矿磁盘加载率（Mb/s）
![`MineWork-LoadingRate.png`](./out/MineWork-LoadingRate_minute_avg.png) 
<!-- - **数据描述**：大部分时间挖矿磁盘加载率为 300Mb/s 左右，表明挖矿阶段磁盘加载率高，占用稳定。 -->

挖矿磁盘加载率是评估系统在挖矿阶段磁盘性能的重要指标。反映了系统在挖矿阶段磁盘的读写速度和效率。挖矿磁盘加载率过低可能导致系统在处理高并发任务时出现延迟，进而影响整体性能。

本次持续性测试期间，采用自动化监控工具对挖矿磁盘加载率进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 17276次挖矿磁盘加载率指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，节点在处理挖矿任务期间，挖矿磁盘加载率均稳定在 300Mb/s 左右，体现出系统在全部时间内都能高效地处理磁盘读写任务。

上图选取了测试期间具有代表性的一天数据，绘制出挖矿磁盘加载率随时间变化的趋势图。从该图可以清晰观察到挖矿磁盘加载率在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在挖矿磁盘加载率方面表现优异，能够高效处理磁盘读写任务，满足实际业务场景中对于响应速度和稳定性的需求。

### 5. 服务器性能指标
![`server_ssd.png`](./out/server/ssd/monitoring.jpg) 
![`server_ssd.png`](./out/server/ssd/resource.jpg) 

服务器性能指标是衡量系统在运行过程中服务器资源使用情况的重要指标。通过对 CPU 使用率、读 IOPS、写 IOPS 和读写磁盘延迟的监控，可以全面了解系统在运行过程中的资源使用情况和性能表现。

本次测试期间，通过自动化监控工具对服务器性能指标进行了实时采集与记录。测试数据显示：

- **CPU 使用率**：大部分时间 CPU 使用率保持在 12.5% 以下，表明系统在运行过程中 CPU 资源使用合理，未出现资源瓶颈。
- **读 IOPS**：平均值 400 以下，峰值 800，表明系统在读取数据时表现出较高的吞吐量。
- **写 IOPS**：平均值 30 以下，峰值 100，表明系统在写入数据时表现出较高的稳定性。
- **读硬盘延迟**：大部分延迟在 1 毫秒以内，极少数情况下超过 2 毫秒，平均延迟 1 毫秒，表明系统在读取数据时表现出较低的延迟。
- **写硬盘延迟**：大部分延迟在 9 毫秒以内，极少数情况下超过 40 毫秒，平均延迟 22 毫秒，表明系统在写入数据时表现出较低的延迟。

以上图表展示了服务器性能指标的折线图，从图中可以清晰观察到各项指标在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在服务器性能指标方面表现优异，能够高效处理高并发任务，满足实际业务场景中对于响应速度和稳定性的需求。

<!-- ## 五、结果分析

通过对图表的分析，我们可以观察到：

- **同步进度差异**：系统在大多数时间内保持了高效的同步状态，确保了数据的一致性和完整性。
- **内存池刷新效率**：内存管理高效，刷新时间短，极大地提高了系统的响应速度。
- **事务同步时间**：事务处理稳定，确保了高并发情况下的系统性能。
- **同步任务积压**：任务调度高效，未完成任务数量少，表明系统在高负载下仍能保持良好的性能。
- **服务器性能指标**：服务器在大多数时间内表现良好，CPU 和内存使用率较低，硬盘读写速度稳定，IOPS 和延迟均在可接受范围内，表明系统在高负载下仍能保持良好的性能。 -->

## 五、结论

经过长达 100 天的持续监测，本次测试结果表明基于区块链技术的分布式存储系统在同步进度差异、内存池刷新效率、事务同步时间、同步任务积压、挖矿磁盘加载率以及服务器性能指标等方面均表现出色。系统能够高效、稳定地处理链上事务和存储任务，满足实际业务场景中对于及时性、稳定性和高效性的需求。

## 六、建议

- 优化内存管理：尽管内存池刷新效率较高，但在极少数情况下仍存在刷新时间超过60微秒的情况，建议进一步优化内存管理策略，减少极端情况下的刷新时间。
- 提升磁盘性能：尽管挖矿磁盘加载率较高，但在高负载情况下仍存在磁盘延迟较高的情况，建议优化磁盘读写策略，提升磁盘性能。
- 增强任务调度能力：尽管同步任务积压较少，但在高负载情况下仍存在未完成任务数量较多的情况，建议增强任务调度能力，确保在高负载情况下的任务处理效率。

## 七、附录
**数据文件**
- 同步进度差：[`SyncProgressDiff.csv`](./out/SyncProgressDiff.csv)
- 内存池刷新效率：[`MemPoolRefreshRate.csv`](./out/MemPoolRefreshRate.csv)
- 事务同步时间：[`TxSyncCompleteTimeCost.csv`](./out/TxSyncCompleteTimeCost.csv)
- 同步任务积压：[`SyncTaskBacklog.csv`](./out/SyncTaskBacklog.csv)
- 挖矿磁盘加载率：[`MineWork.csv`](./out/MineWork.csv)

**部分测试日志**
```log
2025-03-02T00:00:00.015753Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(38.242.218.125:34151), Node: 0x95b4..ff14, addr: 38.242.218.125:48591
2025-03-02T00:00:00.051362Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(157.173.98.166:1234), Node: 0xab6d..1909, addr: 157.173.98.166:1234
2025-03-02T00:00:00.085497Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(217.76.53.34:1234), Node: 0x0c6d..8ce1, addr: 217.76.53.34:1234
2025-03-02T00:00:00.115577Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.167.212.50:34151), Node: 0xe9f3..2e14, addr: 43.167.212.50:24545
2025-03-02T00:00:00.165681Z DEBUG network::discovery: Discovery query completed peers_found=16
2025-03-02T00:00:00.166610Z DEBUG network::peer_manager: Starting a new peer discovery query connected=0 target=50 outbound=0 wanted=16
2025-03-02T00:00:00.166623Z DEBUG network::discovery: Starting a peer discovery request target_peers=16
2025-03-02T00:00:00.204257Z  INFO rpc::zgs::r#impl: zgs_getStatus()
2025-03-02T00:00:00.232246Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.167.173.90:34151), Node: 0xe1e7..abaf, addr: 43.167.173.90:61203
2025-03-02T00:00:00.255637Z  INFO rpc::zgs::r#impl: zgs_getStatus()
2025-03-02T00:00:00.330527Z  INFO rpc::zgs::r#impl: zgs_getStatus()
2025-03-02T00:00:00.342894Z DEBUG sync::controllers::serial: transition started self.tx_seq=3287626 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342906Z DEBUG sync::controllers::serial: transition ended self.tx_seq=3287626 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342909Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253257 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342910Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253257 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342912Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253258 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342913Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253258 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342914Z DEBUG sync::controllers::serial: transition started self.tx_seq=3286443 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342916Z DEBUG sync::controllers::serial: transition ended self.tx_seq=3286443 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342920Z DEBUG sync::controllers::serial: transition started self.tx_seq=3407635 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342922Z DEBUG sync::controllers::serial: transition ended self.tx_seq=3407635 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342923Z DEBUG sync::controllers::serial: transition started self.tx_seq=4443680 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342925Z DEBUG sync::controllers::serial: transition ended self.tx_seq=4443680 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342926Z DEBUG sync::controllers::serial: transition started self.tx_seq=4176634 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342927Z DEBUG sync::controllers::serial: transition ended self.tx_seq=4176634 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342928Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253259 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342929Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253259 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342936Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253260 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342940Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253260 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342943Z DEBUG sync::controllers::serial: transition started self.tx_seq=3021372 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342944Z DEBUG sync::controllers::serial: transition ended self.tx_seq=3021372 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342945Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253262 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342946Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253262 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342947Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253251 self.state=FindingPeers { origin: "5 seconds ago", since: "5 seconds ago" }
2025-03-02T00:00:00.342948Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253251 self.state=Failed { reason: TimeoutFindFile }
2025-03-02T00:00:00.342950Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253261 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342953Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253261 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342954Z DEBUG sync::controllers::serial: transition started self.tx_seq=3950622 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342956Z DEBUG sync::controllers::serial: transition ended self.tx_seq=3950622 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342956Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253256 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342958Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253256 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342959Z DEBUG sync::controllers::serial: transition started self.tx_seq=3179933 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342960Z DEBUG sync::controllers::serial: transition ended self.tx_seq=3179933 self.state=FindingPeers { origin: "2 seconds ago", since: "2 seconds ago" }
2025-03-02T00:00:00.342961Z DEBUG sync::service: Sync stat: incompleted = [3287626, 5253257, 5253258, 3286443, 3407635, 4443680, 4176634, 5253259, 5253260, 3021372, 5253262, 5253251, 5253261, 3950622, 5253256, 3179933], completed = []
2025-03-02T00:00:00.362890Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.167.160.52:34151), Node: 0xe3a7..2df4, addr: 43.167.160.52:36036
2025-03-02T00:00:00.382465Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.153.132.44:34151), Node: 0xdf0d..1837, addr: 43.153.132.44:13688
2025-03-02T00:00:00.384793Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.131.42.33:34151), Node: 0xcdf8..8a48, addr: 43.131.42.33:18971
2025-03-02T00:00:00.397963Z  INFO rpc::zgs::r#impl: zgs_getStatus()
2025-03-02T00:00:00.422698Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.155.212.191:34151), Node: 0x8db8..05ac, addr: 43.155.212.191:59218
2025-03-02T00:00:00.475281Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.166.2.123:34151), Node: 0xe3d0..d6bd, addr: 43.166.2.123:24493
2025-03-02T00:00:00.501652Z  INFO rpc::zgs::r#impl: zgs_getStatus()
2025-03-02T00:00:00.571327Z DEBUG log_entry_sync::sync_manager::log_entry_fetcher: from block number 3451530, latest block number 3451532, confirmation delay 3
2025-03-02T00:00:00.571335Z DEBUG log_entry_sync::sync_manager::log_entry_fetcher: log sync gets entries without progress? old_progress=3451530
2025-03-02T00:00:00.587109Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.133.1.168:34151), Node: 0xf888..642d, addr: 43.133.1.168:28888
2025-03-02T00:00:00.664938Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.163.94.56:34151), Node: 0x10ec..1d10, addr: 43.163.94.56:35632
2025-03-02T00:00:00.870422Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(43.134.73.120:34151), Node: 0x9726..cde6, addr: 43.134.73.120:11437
2025-03-02T00:00:00.953744Z  INFO sync::service: Terminate file sync min_tx_seq=5253257 is_reverted=false
2025-03-02T00:00:00.953755Z DEBUG sync::service: File sync terminated to_terminate=[5253257]
2025-03-02T00:00:00.953759Z  INFO sync::auto_sync::batcher: Terminate file sync due to file already completed in db tx_seq=5253257 num_terminated=1 tx_status=Finalized
2025-03-02T00:00:00.953788Z DEBUG sync::auto_sync::batcher_random: Completed to sync file, state = Ok(RandomBatcherState { name: "random_historical", tasks: [5253251, 5253256, 5253258, 5253259, 5253260, 5253261, 5253262], pending_txs: 2836315, ready_txs: 10, cached_ready_txs: 0 }) tx_seq=5253257 sync_result=Completed
2025-03-02T00:00:00.954056Z DEBUG sync::auto_sync::batcher: Failed to sync file and terminate the failed file sync reason=TimeoutFindFile
2025-03-02T00:00:00.954061Z  INFO sync::service: Terminate file sync min_tx_seq=5253251 is_reverted=false
2025-03-02T00:00:00.954064Z DEBUG sync::service: File sync terminated to_terminate=[5253251]
2025-03-02T00:00:00.954093Z DEBUG sync::auto_sync::batcher_random: Completed to sync file, state = Ok(RandomBatcherState { name: "random_historical", tasks: [5253256, 5253258, 5253259, 5253260, 5253261, 5253262], pending_txs: 2836315, ready_txs: 9, cached_ready_txs: 0 }) tx_seq=5253251 sync_result=Timeout
2025-03-02T00:00:00.954328Z DEBUG sync::auto_sync::batcher_random: Pick a file to sync, state = Ok(RandomBatcherState { name: "random_historical", tasks: [5253253, 5253256, 5253258, 5253259, 5253260, 5253261, 5253262], pending_txs: 2836316, ready_txs: 8, cached_ready_txs: 0 })
2025-03-02T00:00:00.954486Z  INFO sync::service: Terminate file sync min_tx_seq=5253259 is_reverted=false
2025-03-02T00:00:00.954491Z DEBUG sync::service: File sync terminated to_terminate=[5253259]
2025-03-02T00:00:00.954495Z  INFO sync::auto_sync::batcher: Terminate file sync due to file already completed in db tx_seq=5253259 num_terminated=1 tx_status=Finalized
2025-03-02T00:00:00.954521Z DEBUG sync::auto_sync::batcher_random: Completed to sync file, state = Ok(RandomBatcherState { name: "random_historical", tasks: [5253253, 5253256, 5253258, 5253260, 5253261, 5253262], pending_txs: 2836316, ready_txs: 8, cached_ready_txs: 0 }) tx_seq=5253259 sync_result=Completed
2025-03-02T00:00:00.954765Z  INFO sync::service: Start to sync file tx_seq=5253253 maybe_range=None maybe_peer=None
2025-03-02T00:00:00.954940Z DEBUG sync::controllers::serial: transition started self.tx_seq=5253253 self.state=Idle
2025-03-02T00:00:00.954947Z  INFO sync::controllers::serial: Finding peers self.tx_seq=5253253 published=true num_new_peers=0
2025-03-02T00:00:00.954950Z DEBUG sync::controllers::serial: transition ended self.tx_seq=5253253 self.state=FindingPeers { origin: "0 seconds ago", since: "0 seconds ago" }
2025-03-02T00:00:00.954962Z DEBUG router::service: Sending pubsub messages count=1 topics=[AskFile]
2025-03-02T00:00:00.955099Z  WARN network::behaviour: Failed to publish message error=InsufficientPeers topic=AskFile
2025-03-02T00:00:01.201522Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(124.156.208.92:34151), Node: 0x15ad..3aa2, addr: 124.156.208.92:33197
2025-03-02T00:00:01.219981Z DEBUG log_entry_sync::sync_manager::log_entry_fetcher: from block number 3451530, latest block number 3451533, confirmation delay 3
2025-03-02T00:00:01.289273Z  INFO rpc::admin::r#impl: admin_getFileLocation()
2025-03-02T00:00:01.289632Z  INFO rpc::admin::r#impl: admin_findFile(5253262)
2025-03-02T00:00:01.289731Z DEBUG router::service: Sending pubsub messages count=1 topics=[FindFile]
2025-03-02T00:00:01.289860Z  WARN network::behaviour: Failed to publish message error=InsufficientPeers topic=FindFile
2025-03-02T00:00:01.328689Z  WARN discv5::handler: Received an authenticated header without a matching WHOAREYOU request. Node: 0xab78..86f4, addr: 112.249.219.67:1026
2025-03-02T00:00:01.411454Z  INFO rpc::zgs::r#impl: zgs_getStatus()
2025-03-02T00:00:01.427233Z  INFO rpc::zgs::r#impl: zgs_getStatus()
2025-03-02T00:00:01.493750Z  INFO rpc::admin::r#impl: admin_getFileLocation()
2025-03-02T00:00:01.537469Z  INFO rpc::zgs::r#impl: zgs_uploadSegmentsByTxSeq tx_seq=5253260 indices=0
2025-03-02T00:00:01.538040Z DEBUG chunk_pool::mem_pool::chunk_pool_inner: Begin to write segment, root=0x7ce8…85fe, segment_size=52480, segment_index=0
2025-03-02T00:00:01.538852Z DEBUG chunk_pool::mem_pool::chunk_write_control: Succeeded to write segment, root=0x7ce8…85fe, seg_index=0, total_writings=7
2025-03-02T00:00:01.538860Z DEBUG chunk_pool::mem_pool::chunk_pool_inner: Queue to finalize transaction for file 0x7ce8…85fe
2025-03-02T00:00:01.538913Z DEBUG chunk_pool::handler: Received task to finalize transaction id=FileID { root: 0x7ce80f65756ffcab5e481a58172e8e4c03fa5d63f18a204666b07fe865b085fe, tx_id: TxID { seq: 5253260, hash: 0x074474622464aab64cfc7791a5eece2da1f19975a1da60e10dccc0a79feb933d } }
2025-03-02T00:00:01.538953Z DEBUG storage::log_store::log_manager: finalize_tx_with_hash: tx=Transaction { stream_ids: [], data: [], data_merkle_root: 0x7ce80f65756ffcab5e481a58172e8e4c03fa5d63f18a204666b07fe865b085fe, merkle_nodes: [(8, 0xff84d31dda7bc47209f296fc45574f4e78ce51cda6180c367e06a13e4eef9e39), (7, 0x9879e11a2fecbff7e9a7407f4cd20a74a35d8f72ca2031d890f3822bbd2fd458), (5, 0x858693a37291474feb7b97417240e91cefd91f90621d44e1f667d0c9c3878559)], start_entry_index: 9365205120, size: 52466, seq: 5253260 }
2025-03-02T00:00:01.538973Z DEBUG storage::log_store::log_manager: segments_for_proof: 1, last_segment_size_for_proof: 208
2025-03-02T00:00:01.538975Z DEBUG storage::log_store::log_manager: segments_for_file: 1, last_segment_size_for_file: 205
2025-03-02T00:00:01.538976Z DEBUG storage::log_store::log_manager: Padding size: 768
2025-03-02T00:00:01.539498Z DEBUG chunk_pool::handler: Transaction finalized id=FileID { root: 0x7ce80f65756ffcab5e481a58172e8e4c03fa5d63f18a204666b07fe865b085fe, tx_id: TxID { seq: 5253260, hash: 0x074474622464aab64cfc7791a5eece2da1f19975a1da60e10dccc0a79feb933d } } elapsed=577.051µs
2025-03-02T00:00:01.539637Z  WARN network::behaviour: Failed to publish message error=InsufficientPeers topic=NewFile
2025-03-02T00:00:01.539645Z DEBUG router::service: Publish NewFile message new_file=ShardedFile { tx_id: TxID { seq: 5253260, hash: 0x074474622464aab64cfc7791a5eece2da1f19975a1da60e10dccc0a79feb933d }, shard_config: ShardConfig { num_shard: 1, shard_id: 0 } }
2025-03-02T00:00:01.554474Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(217.76.51.115:1234), Node: 0x36a2..b589, addr: 217.76.51.115:1234
2025-03-02T00:00:01.580353Z  WARN discv5::handler: Session has invalid ENR. Enr socket: Some(75.119.159.211:1234), Node: 0x6539..8232, addr: 75.119.159.211:1234
2025-03-02T00:00:01.581964Z  INFO log_entry_sync::sync_manager::log_entry_fetcher: synced 7 events
```
