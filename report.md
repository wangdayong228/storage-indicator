# 数据存储系统性能测试报告

## 一、概述

本报告旨在评估 **[数据存储系统名称]** 在持续高负载下的性能表现。测试模拟了真实生产环境中每天 24 小时、持续 100 天、总计 2400 工时的读写操作，以评估系统的稳定性、吞吐量、延迟等关键指标。

## 二、测试环境

### 1. 硬件环境

- **服务器型号**: [服务器型号]
- **CPU**: [CPU 型号及核数]
- **内存**: [内存容量]
- **存储**: [存储类型及容量，如 SSD、HDD]
- **网络**: [网络带宽]

### 2. 软件环境

- **操作系统**: [操作系统版本]
- **数据存储系统**: [数据存储系统名称及版本]
- **测试工具**: [测试工具名称及版本]

### 3. 测试数据

- **数据类型**: [数据类型，如文本、图片、视频等]

## 三、测试方法

### 1. 测试场景
模拟真实生产环境，设置每天 8-12 小时的读写操作，持续 100 天。

### 2. 测试指标

测试工具: 使用 [测试工具名称] 模拟用户请求，并记录测试指标

**系统指标**
- 同步进度差异
- 内存池刷新效率
- 事务同步时间
- 同步任务积压

**服务器指标**
- 网络吞吐量: 每秒处理的读写请求数量 (IOPS)。
- 读写硬盘延迟: 每个读写请求的平均响应时间。
- 资源利用率: CPU、内存、磁盘、网络等资源的占用情况。

<!-- 本次测试旨在评估基于区块链的分布式存储系统与数据治理平台的性能表现。测试数据来源于系统生成的日志文件，通过一系列的提取器（extractors）将日志数据转换为可分析的CSV格式文件。随后，使用Python脚本生成图表以可视化数据。

#### 数据提取

- 使用Go语言编写的提取器从日志文件中提取关键信息，并生成CSV文件。
- 提取器包括：`SyncProgressDiff`、`MemPoolRefreshRate`、`TxSyncCompleteTimeCost`、`SyncTaskBacklog`、`MineWork`。

#### 数据可视化

- 使用Python的Matplotlib库生成折线图和箱型图。
- 图表包括：同步进度差异、内存池刷新效率、事务同步时间、同步任务积压、挖矿阶段耗时分布。 -->

## 四、测试结果

### 1. 同步进度差异

- 数据文件：[`SyncProgressDiff.csv`](./out/SyncProgressDiff.csv)
- 图表：箱型图显示了不同时间点的区块同步进度差异。
![`SyncProgressDiff.png`](./out/SyncProgressDiff_boxplot.png) 
- **数据描述**：大部分同步进度差异集中在2个区块以内，表明系统在大多数时间内保持了高效的同步状态。

### 2. 内存池刷新效率

- 数据文件：[`MemPoolRefreshRate.csv`](./out/MemPoolRefreshRate.csv)
- 图表：箱型图展示了内存池刷新时间的分布情况。
![`MemPoolRefreshRate.png`](./out/MemPoolRefreshRate_boxplot.png) 
- **数据描述**：大部分刷新时间集中在10微秒以内，极少数情况下超过60微秒，显示出系统在内存管理上的高效性。

### 3. 事务同步时间

- 数据文件：[`TxSyncCompleteTimeCost.csv`](./out/TxSyncCompleteTimeCost.csv)
- 图表：箱型图展示了事务同步完成所需时间的分布。
![`TxSyncCompleteTimeCost.png`](./out/TxSyncCompleteTimeCost_boxplot.png) 
- **数据描述**：大部分事务同步时间在10秒以内，显示出系统在事务处理上的稳定性。

### 4. 同步任务积压

- 数据文件：[`SyncTaskBacklog.csv`](./out/SyncTaskBacklog.csv)
- 图表：箱型图显示了不同时间点的未完成任务数量。
![`SyncTaskBacklog.png`](./out/SyncTaskBacklog_boxplot.png) 
- **数据描述**：未完成任务数量大多保持在16个以下，表明系统在任务调度上的高效性。

### 5. 挖矿磁盘加载率（Mb/s）

- 数据文件：[`MineWork.csv`](./out/MineWork.csv)
- 图表：折线图显示了不同时间点的挖矿磁盘加载率。
![`MineWork-LoadingRate.png`](./out/MineWork-LoadingRate_minute_avg.png) 
- **数据描述**：大部分时间挖矿磁盘加载率为 300Mb/s 左右，表明挖矿阶段磁盘加载率高，占用稳定。

### 5. 服务器性能指标

- 图表：包括 CPU 使用率、读 IOPS、写 IOPS 和读写磁盘延迟的折线图。
![`server_ssd.png`](./out/server/ssd/monitoring.jpg) 
![`server_ssd.png`](./out/server/ssd/resource.jpg) 
- **数据描述**：
  - **CPU 使用率**：大部分时间 CPU 使用率保持在 12.5% 以下。
  - **读 IOPS**：平均值 400 以下，峰值800。
  - **写 IOPS**：平均值 30 以下，峰值100。
  - **读硬盘延迟**：大部分延迟在 1 毫秒以内，极少数情况下超过 2 毫秒，平均延迟 1 毫秒。
  - **写硬盘延迟**：大部分延迟在 9 毫秒以内，极少数情况下超过 40 毫秒，平均延迟 22 毫秒。

## 五、结果分析

通过对图表的分析，我们可以观察到：

- **同步进度差异**：系统在大多数时间内保持了高效的同步状态，确保了数据的一致性和完整性。
- **内存池刷新效率**：内存管理高效，刷新时间短，极大地提高了系统的响应速度。
- **事务同步时间**：事务处理稳定，确保了高并发情况下的系统性能。
- **同步任务积压**：任务调度高效，未完成任务数量少，表明系统在高负载下仍能保持良好的性能。
- **服务器性能指标**：服务器在大多数时间内表现良好，CPU 和内存使用率较低，硬盘读写速度稳定，IOPS 和延迟均在可接受范围内，表明系统在高负载下仍能保持良好的性能。

## 六、结论

本次测试表明，基于区块链的分布式存储系统在正常负载下表现优异，尤其在数据同步、内存管理和任务调度方面展现了显著的优势。服务器性能指标显示，系统在高负载情况下仍能保持良好的性能，未来的优化可以进一步提升系统在极端条件下的表现。

## 七、建议

- 建议在实际生产环境中部署 **[数据存储系统名称]**，并进行进一步的性能优化和监控。
- 建议定期进行性能测试，以评估系统性能变化趋势，并及时发现潜在问题。

## 八、附录
- **部分测试日志**

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
