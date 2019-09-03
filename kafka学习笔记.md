# kafka 学习笔记

## 1. 基本概念

### 1.1 Topic, Partition

> 一个 Topic 就对应一个 消息队列（MQ），为利于横向扩展，以便能处理大数据，kafka 将一个 Topic 分成多个 Partition，存储的时候，消息会平均分配到一个 Partition。Partition 中会为消息保存一个 Partition 内唯一的 ID ，一般称为偏移量(offset)。

### 1.2 消费模型

> 一般有两种消费模型：队列模式（点对点模式）、发布/订阅模式

* 队列模式（也叫点对点模式）。多个消费者共同消费一个队列，每条消息只发送给一个消费者。

* 发布/订阅模式。多个消费者订阅主题，每个消息会发布给所有的消费者。

### 1.3 Partition 与消费模型

### 1.4 物理存储

> Topic, Partition 都是抽象的概念。每个 Partition 最终都需要存储在物理机器上，在 Kafka 中一般把这样的物理机器称为 `Broker`，可以是一台物理机，也可以是一个集群。

### 1.5 消费的 offset

> **auto.offset.reset** 
>
> ​	值为 latest 则从最近一个开始消费（含上次落在的最近的一个 offset）
>
> ​	值为 smallest 则是从最开始的 offset 开始消费，类似命令行的 from-beginning

### 1.6 kafka 配置多个 broker 

[参考文档](<https://www.cnblogs.com/yingww/p/8746379.html>)

> 
>
>

​    