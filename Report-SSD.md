<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [数据存储系统性能测试报告](#%E6%95%B0%E6%8D%AE%E5%AD%98%E5%82%A8%E7%B3%BB%E7%BB%9F%E6%80%A7%E8%83%BD%E6%B5%8B%E8%AF%95%E6%8A%A5%E5%91%8A)
  - [一、概述](#%E4%B8%80%E6%A6%82%E8%BF%B0)
  - [二、测试环境](#%E4%BA%8C%E6%B5%8B%E8%AF%95%E7%8E%AF%E5%A2%83)
    - [1. 硬件环境](#1-%E7%A1%AC%E4%BB%B6%E7%8E%AF%E5%A2%83)
    - [2. 软件环境](#2-%E8%BD%AF%E4%BB%B6%E7%8E%AF%E5%A2%83)
  - [三、测试方法](#%E4%B8%89%E6%B5%8B%E8%AF%95%E6%96%B9%E6%B3%95)
    - [1. 测试场景描述](#1-%E6%B5%8B%E8%AF%95%E5%9C%BA%E6%99%AF%E6%8F%8F%E8%BF%B0)
    - [2. 测试过程和监控方式](#2-%E6%B5%8B%E8%AF%95%E8%BF%87%E7%A8%8B%E5%92%8C%E7%9B%91%E6%8E%A7%E6%96%B9%E5%BC%8F)
  - [四、测试结果](#%E5%9B%9B%E6%B5%8B%E8%AF%95%E7%BB%93%E6%9E%9C)
    - [1. 同步进度差](#1-%E5%90%8C%E6%AD%A5%E8%BF%9B%E5%BA%A6%E5%B7%AE)
    - [2. 内存池刷新效率](#2-%E5%86%85%E5%AD%98%E6%B1%A0%E5%88%B7%E6%96%B0%E6%95%88%E7%8E%87)
    - [3. 事务同步时间](#3-%E4%BA%8B%E5%8A%A1%E5%90%8C%E6%AD%A5%E6%97%B6%E9%97%B4)
    - [4. 同步任务积压](#4-%E5%90%8C%E6%AD%A5%E4%BB%BB%E5%8A%A1%E7%A7%AF%E5%8E%8B)
    - [5. 挖矿磁盘加载率（Mb/s）](#5-%E6%8C%96%E7%9F%BF%E7%A3%81%E7%9B%98%E5%8A%A0%E8%BD%BD%E7%8E%87mbs)
    - [5. 服务器性能指标](#5-%E6%9C%8D%E5%8A%A1%E5%99%A8%E6%80%A7%E8%83%BD%E6%8C%87%E6%A0%87)
  - [五、测试结果解析](#%E4%BA%94%E6%B5%8B%E8%AF%95%E7%BB%93%E6%9E%9C%E8%A7%A3%E6%9E%90)
    - [（一）系统稳定性表现优异](#%E4%B8%80%E7%B3%BB%E7%BB%9F%E7%A8%B3%E5%AE%9A%E6%80%A7%E8%A1%A8%E7%8E%B0%E4%BC%98%E5%BC%82)
    - [（二）高效的内存管理机制](#%E4%BA%8C%E9%AB%98%E6%95%88%E7%9A%84%E5%86%85%E5%AD%98%E7%AE%A1%E7%90%86%E6%9C%BA%E5%88%B6)
    - [（三）任务调度与处理能力卓越](#%E4%B8%89%E4%BB%BB%E5%8A%A1%E8%B0%83%E5%BA%A6%E4%B8%8E%E5%A4%84%E7%90%86%E8%83%BD%E5%8A%9B%E5%8D%93%E8%B6%8A)
    - [（四）磁盘性能满足业务高负载需求](#%E5%9B%9B%E7%A3%81%E7%9B%98%E6%80%A7%E8%83%BD%E6%BB%A1%E8%B6%B3%E4%B8%9A%E5%8A%A1%E9%AB%98%E8%B4%9F%E8%BD%BD%E9%9C%80%E6%B1%82)
    - [（五）服务器资源利用高效合理](#%E4%BA%94%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%B5%84%E6%BA%90%E5%88%A9%E7%94%A8%E9%AB%98%E6%95%88%E5%90%88%E7%90%86)
  - [六、测试结论](#%E5%85%AD%E6%B5%8B%E8%AF%95%E7%BB%93%E8%AE%BA)
  - [七、附录](#%E4%B8%83%E9%99%84%E5%BD%95)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->



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
![`SyncProgressDiff.png`](./out/ssd/SyncProgressDiff_boxplot.png) 

同步进度差是衡量基于区块链的分布式存储系统节点与区块链主链同步进度差异的重要技术指标。它反映了节点能否及时有效地跟踪区块链网络上的最新状态，并实时处理链上订单和交易任务的能力。同步进度差异过大可能导致节点无法及时获取链上数据，进而影响系统的一致性、稳定性和业务连续性。

本次持续性测试期间，采用自动化监控工具对节点同步进度差异进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 124718 次同步进度差异指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，无论节点处于持续读写负载还是空闲状态下，每小时同步进度差异的 75% 分位数均稳定在 2 个区块以内，体现出节点在大部分时间内都能高效地追踪主链最新状态。此外，在整个测试过程中，节点所记录的同步进度差异最大值也未超过 4 个区块，这表明即使在高负载或网络波动的情况下，节点仍能较快跟进至最新区块高度。

上图选取了测试期间具有代表性的一天数据，绘制出节点同步进度差异随时间变化的趋势图。从该图可以清晰观察到同步进度差异在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明节点与区块链主链之间的同步进度差异始终保持在较低水平，节点能够可靠、高效地追踪链上状态，满足实际业务场景中对于及时性和稳定性的需求。

### 2. 内存池刷新效率
![`MemPoolRefreshRate.png`](./out/ssd/MemPoolRefreshRate_boxplot.png) 

内存池刷新效率是衡量系统在内存管理方面性能的重要指标。反映了系统在处理内存数据时的响应速度和效率。内存池刷新效率低下可能导致系统在处理高并发任务时出现延迟，进而影响整体性能。

本次持续性测试期间，采用自动化监控工具对内存池刷新时间进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 51844 次内存池刷新时间指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，无论节点处于持续读写负载还是空闲状态下，内存池刷新时间的 75% 分位数均稳定在 10 微秒以内，体现出系统在大部分时间内都能高效地处理内存数据。此外，在整个测试过程中，内存池刷新时间的最大值也未超过 60 微秒，这表明即使在高负载或网络波动的情况下，系统仍能保持较高的内存管理效率。

上图选取了测试期间具有代表性的一天数据，绘制出内存池刷新时间随时间变化的趋势图。从该图可以清晰观察到内存池刷新时间在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在内存池刷新效率方面表现优异，能够高效处理内存数据，满足实际业务场景中对于响应速度和稳定性的需求。

### 3. 事务同步时间

![`TxSyncCompleteTimeCost.png`](./out/ssd/TxSyncCompleteTimeCost_boxplot.png) 

事务同步时间是评估系统在处理链上事务时的性能指标。反映了系统在处理链上事务时的响应速度和效率。事务同步时间过长可能导致系统在处理高并发任务时出现延迟，进而影响整体性能。

本次持续性测试期间，采用自动化监控工具对事务同步时间进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 4910 次事务同步时间指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，无论节点处于持续读写负载还是空闲状态下，事务同步时间的 75% 分位数均稳定在 10 秒以内，体现出系统在大部分时间内都能高效地处理链上事务。<!-- 此外，在整个测试过程中，事务同步时间的最大值也未超过 20 秒，这表明即使在高负载或网络波动的情况下，系统仍能保持较高的事务处理效率。-->


上图选取了测试期间具有代表性的一天数据，绘制出事务同步时间随时间变化的趋势图。从该图可以清晰观察到事务同步时间在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在事务同步时间方面表现优异，能够高效处理链上事务，满足实际业务场景中对于响应速度和稳定性的需求。


### 4. 同步任务积压

![`SyncTaskBacklog.png`](./out/ssd/SyncTaskBacklog_boxplot.png) 

同步任务积压是衡量系统在任务调度方面性能的重要指标。反映了系统在处理任务时的响应速度和效率。同步任务积压过多可能导致系统在处理高并发任务时出现延迟，进而影响整体性能。

本次持续性测试期间，采用自动化监控工具对同步任务积压进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 28801 次同步任务积压指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，无论节点处于持续读写负载还是空闲状态下，同步任务积压的数量 100% 分位数均稳定在 16 个以下，体现出系统在全部时间内都能高效地处理任务。

上图选取了测试期间具有代表性的一天数据，绘制出同步任务积压随时间变化的趋势图。从该图可以清晰观察到同步任务积压在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在同步任务积压方面表现优异，能够高效处理任务，满足实际业务场景中对于响应速度和稳定性的需求。

### 5. 挖矿磁盘加载率（Mb/s）
![`MineWork-LoadingRate.png`](./out/ssd/MineWork-LoadingRate_minute_avg.png) 
<!-- - **数据描述**：大部分时间挖矿磁盘加载率为 300Mb/s 左右，表明挖矿阶段磁盘加载率高，占用稳定。 -->

挖矿磁盘加载率是评估系统在挖矿阶段磁盘性能的重要指标。反映了系统在挖矿阶段磁盘的读写速度和效率。挖矿磁盘加载率过低可能导致系统在处理高并发任务时出现延迟，进而影响整体性能。

本次持续性测试期间，采用自动化监控工具对挖矿磁盘加载率进行了长期不间断的实时采集与记录，在测试过程中共累计采集了 17276次挖矿磁盘加载率指标数据，涵盖了系统处于不同负载状态下的表现，包括高负载情形（即读写压测阶段）以及相对空闲的低负载情形。

统计分析显示，节点在处理挖矿任务期间，挖矿磁盘加载率均稳定在 300Mb/s 左右，体现出系统在全部时间内都能高效地处理磁盘读写任务。

上图选取了测试期间具有代表性的一天数据，绘制出挖矿磁盘加载率随时间变化的趋势图。从该图可以清晰观察到挖矿磁盘加载率在全天范围内的波动情况及趋势变化。

总体而言，经过长达 100 天的持续监测，本次测试结果表明系统在挖矿磁盘加载率方面表现优异，能够高效处理磁盘读写任务，满足实际业务场景中对于响应速度和稳定性的需求。

### 5. 服务器性能指标
![`server_ssd.png`](./out/server/ssd/zg-ssd-monitor-cut.png) 
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

## 五、测试结果解析

本次对基于区块链技术的分布式存储系统开展的长期持续性性能测试，历时 100 天，总计约 2400 小时，全面、系统地评估了系统在长期高强度负载条件下的稳定性、可靠性及性能表现。通过对同步进度差、内存池刷新效率、事务同步时间、同步任务积压、挖矿磁盘加载率以及服务器资源使用情况等关键性能指标的实时监控与深度分析，我们得出以下结论：

### （一）系统稳定性表现优异

测试过程中，节点与区块链主链的同步进度差异始终维持在较低水平，75% 分位数稳定于 2 个区块以内，最大值未超过 4 个区块。这表明系统节点能够长期稳定地追踪区块链主链的最新状态，确保了数据与任务处理的及时性以及整体业务的一致性。即使在高负载情境以及网络波动情况下，系统仍能快速完成与主链的同步，彰显了良好的稳定性与健壮性。

此外，系统在事务同步方面表现同样出色，事务同步时间 75% 分位数稳定在 10 秒以内，表明链上事务处理的高效性与稳定性。整体来看，系统在各种复杂条件与负载强度下，均能保持数据同步与事务处理的稳定性，充分满足实际业务场景对长期稳定运行的需求。

### （二）高效的内存管理机制

内存池刷新效率是系统内存管理性能的重要体现，通过对 51844 次内存池刷新时间的监测与分析，我们发现在长期测试过程中系统的内存池刷新效率表现杰出。统计数据显示 75% 分位数均稳定在 10 微秒以内，最大值未超过 60 微秒，显示出系统具备高度高效的内存数据处理能力。即便是在高并发读写请求的情境下，系统内存管理机制仍然能够快速反应，保证了系统资源的高效利用，避免了因内存性能不足而产生的性能瓶颈。

### （三）任务调度与处理能力卓越

同步任务积压指标体现了系统在任务调度与处理方面的性能。测试期间，总计 28801 次的同步任务积压监测数据表明，系统在所有负载条件下积压任务数量始终维持在较低水平（最大值小于 16 个任务）。这意味着即使在长期高负载运行环境下，系统也能通过有效的任务调度策略，及时高效地处理任务队列，避免任务积压过多引发的延迟和性能下降。此指标进一步说明了系统在长期高负载场景下的优异任务处理能力及其对高并发业务场景的适应性。

### （四）磁盘性能满足业务高负载需求

挖矿磁盘加载率指标反映了系统在挖矿业务场景下的磁盘读写能力。测试期间，对 17276 次挖矿磁盘加载率的监测显示，系统磁盘加载率在挖矿任务运行中持续稳定在 300Mb/s 左右，表现出稳定而高效的磁盘读写性能。这种持续稳定的磁盘性能，能够有效支撑高并发的读写操作，确保系统在长期挖矿任务以及数据存储业务中表现出色，避免磁盘性能瓶颈问题，保证了系统整体性能的持续稳定和高效运行。

### （五）服务器资源利用高效合理

服务器性能指标是理解系统整体资源利用效率的重要依据。测试期间，服务器 CPU 使用率整体维持在较低水平（12.5% 以下），有效避免了 CPU 资源瓶颈问题的出现；磁盘的读写 IOPS 表现合理，读 IOPS 平均值维持在 400 以下、峰值 800，写 IOPS 平均值 30 以下、峰值 100，均体现了系统读写操作的高效与稳定；磁盘读写延迟指标也表现出色，读延迟平均 1 毫秒，写延迟平均 22 毫秒，整体延迟维持在较低水平，进一步保障了系统数据访问的快速响应。

总体来看，服务器资源的利用情况表明系统设计合理，资源调度高效，能够在长期运行环境中提供稳定而高效的性能表现，有效满足生产环境中实际业务场景的需求。

## 六、测试结论

综合以上关键性能指标的监测与分析结果，我们可以明确得出，本次送测的基于区块链技术的分布式存储系统在长达 100 天的持续高强度负载条件下，整体表现优异，稳定性、可靠性及性能均达到了预期目标。系统在同步效率、内存管理、任务调度、磁盘读写性能以及服务器资源利用上均表现出了显著的优势，能够稳定、高效地满足实际业务场景的长期运行需求。

综上所述，本次测试验证了系统设计的合理性与技术路线的可行性，系统整体性能表现完全符合预期，具备上线实际生产环境的性能条件。同时，测试过程中未发现明显性能瓶颈或异常情况，进一步证实了系统的稳健性与高效性。我们认为，该分布式存储系统在长期运行高并发负载下，具备良好的适应性、可靠性和稳定性，能够为实际业务运营提供持续稳定的性能保障。

## 七、附录
**数据文件**
- 同步进度差：[`SyncProgressDiff.csv`](./out/ssd/SyncProgressDiff.csv)
- 内存池刷新效率：[`MemPoolRefreshRate.csv`](./out/ssd/MemPoolRefreshRate.csv)
- 事务同步时间：[`TxSyncCompleteTimeCost.csv`](./out/ssd/TxSyncCompleteTimeCost.csv)
- 同步任务积压：[`SyncTaskBacklog.csv`](./out/ssd/SyncTaskBacklog.csv)
- 挖矿磁盘加载率：[`MineWork.csv`](./out/ssd/MineWork.csv)

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
```
