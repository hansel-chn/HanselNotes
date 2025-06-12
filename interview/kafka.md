## 流量上

kafka瓶颈一般在网络流量上，先考虑流量的需求，评估每秒需要处理的数据量，根据需求的

## 存储

根据实际的qps，每条信息的大小，留存时间，估算需要的存储空间，再留一些缓冲的量，

新增消息数
消息留存时间
平均消息大小
备份数
是否启用压缩

## 理论

最多一次（at most once）：消息可能会丢失，但绝不会被重复发送。
至少一次（at least once）：消息不会丢失，但有可能被重复发送。
精确一次（exactly once）：消息不会丢失，也不会被重复发送。

### 幂等性怎么保证的

Producer ID (PID): 每个 Kafka 生产者实例在初始化时会被分配一个唯一的 Producer ID，用于标识该生产者。  
Sequence Number: 每个分区的消息都会维护一个递增的序列号。生产者在发送消息时会附带该序列号，Broker
会根据序列号判断消息是否是重复的。  
Broker 的去重逻辑: Broker 会为每个分区的每个生产者维护一个缓存，用于存储最近的 Producer ID
和对应的序列号。如果收到的消息序列号小于或等于缓存中的序列号，则认为是重复消息并丢弃。  
幂等性配置: 在 Kafka 生产者中，通过设置 enable.idempotence=true 开启幂等性。开启后，Kafka 会自动处理消息的去重逻辑，无需额外的代码实现。

> 以前的版本的kafka生产者，并不能保证幂等性，在后来，加入了生产者幂等性的配置，实际原理是使用额外的空间，在broker保存额外的信息，通过比对生产者ID和消息的序列号，判断消息是否重复

### 多分区消息的幂等性

通过kafka提供的分布式的事务来实现的，通过类似分布式事务的事务的初始化，事务的开始，提交以及中止，来实现消息的精确一次语义

### 为什么要设计消费者组

常见的消息队列模型有点对点的模型和发布订阅的模型，点对点的模型指指定的消息只能且仅仅只能被一个消费者消费。
而发布订阅模型指指定的消息被发布后，可以同时被多个订阅者消费，且互相之间并不受影响，消费者组的设计同时满足了这两种情况，如果消费者属于同一个消费者组则实现了点对点的模型，消费者属于不同的消费者组，则实现了发布订阅的模型

消费者组会基于其订阅的主题分区数,分配给消费者合适的分区数量。这也导致kafka中比较令人头疼的重平衡问题

### kafka是拉取信息

### 重平衡 （何时引发）

* 加入新的消费者或者消费组中的消费者宕机
* 分区数发生变化
* 订阅了新的主题

Rebalanced, Coordinator
消费者组对应的Coordinator，根据消费者组经过hash等运算得到位移主题的哪一个分区保存消费者组的元信息，确定该分区leader对应的broker即为协调者

后两者可以人为的进行控制（不必要的重平衡 ）

解决：消费时间过长带来的消费者离开消费组引发的重平衡,
心跳未收到导致的认为消费者宕机，通过设计心跳频率和超时时间来优化

### 位移主题

kafka消费者的唯一信息，同样也是kafka的一个主题，但是是由kafka自身维护的。
面对自动提交位移，kafka还存在compact机制，压缩冗余的信息，比如无新消息写入时，重复消息的提交

### 多协程消费消费者

与方案1的粗粒度不同，方案2将任务切分成了消息获取和消息处理两个部分，分别由不同的线程处理它们。比起方案1，方案2的最大优势就在于它的高伸缩性，就是说我们可以独立地调节消息获取的线程数，以及消息处理的线程数，而不必考虑两者之间是否相互影响。如果你的消费获取速度慢，那么增加消费获取的线程数即可；如果是消息的处理速度慢，那么增加Worker线程池线程数即可。

## ISR

Kafka 的 Producer acks 和 ISR（In-Sync Replicas）机制是协调消息可靠性和性能的核心部分。以下是它们的协调机制：

还有高水位相关内容（副本同步）

1. Producer acks 参数
   acks 参数决定了生产者在发送消息时需要等待多少副本确认消息写入成功：  
   acks=0: Producer 不等待任何确认，消息可能丢失，但性能最高。
   acks=1: Producer 只等待 Leader 副本确认写入成功，性能较高，但如果 Leader 宕机，可能导致数据丢失。
   acks=all 或 acks=-1: Producer 等待所有 ISR 副本确认写入成功，可靠性最高，但性能较低。
2. ISR（In-Sync Replicas）
   ISR 是 Kafka 中一个动态维护的副本集合，包含所有与 Leader 副本保持同步的副本。只有 ISR 中的副本会参与消息的确认流程。  
   Leader 副本：负责处理 Producer 的写入请求。
   Follower 副本：从 Leader 拉取数据，保持与 Leader 的同步。
3. 协调机制
   当 Producer 发送消息时，以下是 acks 和 ISR 的协调过程：  
   消息写入 Leader：  
   Producer 将消息发送到 Leader 副本。
   Leader 副本将消息写入本地日志，并将消息推送到 ISR 中的 Follower 副本。
   等待确认：  
   如果 acks=0，Producer 不等待任何确认，直接返回。
   如果 acks=1，Producer 等待 Leader 副本写入成功后返回。
   如果 acks=all，Producer 等待 ISR 中所有副本都确认写入成功后返回。
   同步检查：  
   Kafka 定期检查 ISR 中的副本是否与 Leader 同步。
   如果某个 Follower 副本落后太多（超过 replica.lag.time.max.ms），它会被移出 ISR。
   故障处理：
   如果 Leader 副本宕机，Kafka 会从 ISR 中选举一个新的 Leader。
   如果 ISR 中没有副本可用，可能会导致数据丢失（取决于 min.insync.replicas 配置）。
4. min.insync.replicas 参数
   配合 acks=all 使用，确保至少有一定数量的副本（包括 Leader）确认写入成功。
   如果 ISR 中的副本数小于 min.insync.replicas，Kafka 会拒绝写入请求，防止数据丢失。
   总结
   acks 决定了 Producer 等待的确认级别。
   ISR 确保了副本之间的同步状态。
   两者通过 Leader 副本协调，结合 min.insync.replicas 提供可靠性和性能的平衡。

https://github.com/h2pl/JavaTutorial/blob/master/docs/mq/kafka/%E6%B6%88%E6%81%AF%E9%98%9F%E5%88%97kafka%E8%AF%A6%E8%A7%A3%EF%BC%9AKafka%E5%8E%9F%E7%90%86%E5%88%86%E6%9E%90%E6%80%BB%E7%BB%93%E7%AF%87.md

## kafka的零拷贝

重点：https://zhuanlan.zhihu.com/p/616105519  上下文的切换和内存的copy带来的问题

> mmap + write
> 通过使用 mmap() 来代替 read()， 可以减少一次数据拷贝的过程。
> 但这还不是最理想的零拷贝，因为仍然需要通过 CPU 把内核缓冲区的数据拷贝到 socket 缓冲区里，而且仍然需要 4 次上下文切换，因为系统调用还是
> 2 次。

> sendfile
> 于是，从 Linux 内核 2.4 版本开始起，对于支持网卡支持 SG-DMA 技术的情况下， sendfile() 系统调用的过程发生了点变化，具体过程如下：
> 第一步，通过 DMA 将磁盘上的数据拷贝到内核缓冲区里；
> 第二步，缓冲区描述符和数据长度传到 socket 缓冲区，这样网卡的 SG-DMA 控制器就可以直接将内核缓存中的数据拷贝到网卡的缓冲区里，此过程不需要将数据从操作系统内核缓冲区拷贝到
> socket 缓冲区中，这样就减少了一次数据拷贝；

sendfile减少了文件拷贝次数和上下文的切换

Linux 2.4+ 内核通过 sendfile 系统调用，提供了零拷贝。磁盘数据通过 DMA(Direct Memory Access) 拷贝到内核态 Buffer 后，直接通过
DMA 拷贝到 NIC Buffer(socket buffer)，无需 CPU 拷贝。

## ISR

Replication: Messages are replicated across multiple brokers. If a broker fails, replicas ensure data availability.
acks Configuration:
acks=all: The producer waits for all in-sync replicas (ISR) to acknowledge the message before considering it
successfully sent.
Combined with min.insync.replicas, this ensures a minimum number of replicas acknowledge the message.

(
这个标准就是Broker端参数replica.lag.time.max.ms参数值。这个参数的含义是Follower副本能够落后Leader副本的最长时间间隔，当前默认值是10秒。这就是说，只要一个Follower副本落后Leader副本的时间不连续超过10秒，那么Kafka就认为该Follower副本与Leader是同步的，即使此时Follower副本中保存的消息明显少于Leader副本中的消息。
我们在前面说过，Follower副本唯一的工作就是不断地从Leader副本拉取消息，然后写入到自己的提交日志中。如果这个同步过程的速度持续慢于Leader副本的消息写入速度，那么在replica.lag.time.max.ms时间后，此Follower副本就会被认为是与Leader副本不同步的，因此不能再放入ISR中。此时，Kafka会自动收缩ISR集合，将该副本“踢出”ISR。)

## kafka能设置面向不同topic采用不同的消息投递语义吗，从而达到一些允许消息丢失一些允许消息丢失吗

Yes, Kafka can be configured to use different delivery semantics (at-most-once, at-least-once, or exactly-once) for
different topics. This can be achieved by configuring the producer and consumer settings on a per-topic basis. Here's
how:

### 1. **Producer Configuration**

The producer's delivery semantics depend on the following settings:

- **`acks`**: Controls the acknowledgment behavior.
    - `acks=0`: At-most-once (messages may be lost).
    - `acks=1`: At-least-once (default, messages may be lost if the leader fails after acknowledgment but before
      replication).
    - `acks=all`: At-least-once with stronger durability guarantees (requires all replicas to acknowledge).
- **`retries`**: Number of retries for failed sends. Set to `0` for at-most-once semantics.
- **`enable.idempotence`**: Set to `true` for exactly-once semantics (requires `acks=all`).

You can configure these settings differently for each producer instance based on the topic.

### 2. **Consumer Configuration**

The consumer's delivery semantics depend on:

- **`enable.auto.commit`**: Controls whether offsets are committed automatically.
    - `true`: At-least-once (default, offsets are committed automatically after processing).
    - `false`: At-most-once or exactly-once (manual offset management is required).
- **Manual Offset Commit**: For exactly-once, commit offsets only after processing is complete.

### Example: Configuring Different Semantics for Topics

Below is an example in Go using the `sarama` Kafka client library:

```go
package main

import (
	"log"
	"github.com/Shopify/sarama"
)

func createProducer(topic string, allowLoss bool) sarama.SyncProducer {
	config := sarama.NewConfig()
	if allowLoss {
		// At-most-once semantics
		config.Producer.RequiredAcks = sarama.NoResponse
		config.Producer.Retry.Max = 0
	} else {
		// At-least-once semantics
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = 5
		config.Producer.Idempotent = true // For exactly-once
	}

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	return producer
}

func main() {
	// Example: Topic1 allows message loss, Topic2 does not
	producer1 := createProducer("Topic1", true)
	producer2 := createProducer("Topic2", false)

	defer producer1.Close()
	defer producer2.Close()

	// Send messages to Topic1 and Topic2
	// ...
}
```

### Summary

By configuring the producer and consumer settings differently for each topic, you can achieve different delivery
semantics for different topics in Kafka.

## 怎么知道kafka消费不及时积累数据

要判断 Kafka 消费是否不及时并导致数据积累，可以通过以下方法：

### 1. **监控 Consumer Lag**

Consumer Lag 是衡量 Kafka 消费是否及时的关键指标。它表示生产者写入的最新消息与消费者已消费的最新消息之间的差距。

#### 方法：

- 使用 Kafka 自带工具 `kafka-consumer-groups.sh` 查看 Consumer Lag：
  ```bash
  kafka-consumer-groups.sh --bootstrap-server <broker> --group <consumer-group> --describe
  ```
  输出中 `LAG` 列表示未消费的消息数量。如果该值持续增长，说明消费不及时。

- 使用监控工具（如 Prometheus + Grafana）监控 `consumer_lag` 指标。

---

### 2. **监控 Topic 的 Log End Offset**

Log End Offset 是 Kafka 中每个分区的最新消息偏移量。如果消费者的 `Committed Offset` 长期落后于 `Log End Offset`，说明消费不及时。

#### 方法：

- 使用 Kafka 自带工具 `kafka-run-class.sh` 查看分区的 Log End Offset：
  ```bash
  kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list <broker> --topic <topic>
  ```
  比较 Log End Offset 和消费者的 Committed Offset。

---

### 3. **监控 Broker 的消息堆积**

如果消费不及时，Broker 的消息堆积会增加，可能导致磁盘使用率上升。

#### 方法：

- 查看 Broker 的磁盘使用情况，监控 `log.dirs` 中的磁盘空间。
- 监控 Kafka 指标 `UnderReplicatedPartitions` 和 `LogEndOffset`.

---

### 4. **分析消费者性能**

- 检查消费者是否有足够的线程或实例来处理分区。
- 检查消费者是否存在性能瓶颈（如消息处理逻辑过慢）。
- 确保消费者的 `max.poll.records` 和 `fetch.max.bytes` 配置合理。

---

### 5. **使用 Kafka Manager 或其他工具**

使用 Kafka Manager、Confluent Control Center 或其他 Kafka 管理工具，可以直观地查看 Consumer Lag 和消息堆积情况。

通过以上方法，可以有效判断 Kafka 消费是否不及时并导致数据积累。