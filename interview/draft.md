## 一致性hash

> 一致性hash一致性体现在哪里，为什么叫一致性。

因为常见的hash算法是对节点数进行hash运算，这样在节点宕机后，hash计算值变化，所有节点和数据要重新进行数据迁移。
一致性hash对固定值取模，只会影响个别节点，其他节点不受影响。
（Consistent hashing is named for its ability to minimize the number of keys that need to be remapped when the hash table
is resized (e.g., when nodes are added or removed)）

> 一致性hash带来的请求分配不均的问题怎么处理。

构建某一个节点对应的多个虚拟节点，随机均匀分布到环上，使得某一个节点宕机不会完全使单个节点承担所有压力

## redis vs mysql vs elasticsearch

* redis 键值型数据库
  Advantages:
  Extremely fast (in-memory operations).
  Supports data structures like lists, sets, and sorted sets.
  Persistence options (RDB, AOF) for durability.
  Limitations:
  Limited by memory size.
  Not ideal for complex queries or large datasets.

* mysql 关系型数据库
  结构化的数据，良好的事务支持，
* elasticsearch
  全文本搜索，日志分析，近实时的搜索和分析，处理结构化和半结构化数据，支持很多的复杂查询，

Redis: For caching, real-time data, or ephemeral data storage.
MySQL: For structured, relational data with strict consistency requirements.
Elasticsearch: For search-heavy applications or analytics on large datasets.

## http2优势

* 以前的问题，头部巨大，对头部进行了压缩，节省了带宽，（经过霍夫曼编码的静态动态表）减少字符长度
* http队头阻塞的问题，并发的6个链接，同一链接在处理完一个http事务后，才能处理下一个事务。现在每一个http连接可以有多个stream(
  可以乱序发送)， 并且一个连接中多个stream的存在允许多个请求相应并发发送,每个stream还是专注于单个请求响应对。
* 支持服务器推送资源

Frames are the smallest unit of communication in HTTP/2. They carry raw data but lack semantic meaning.
Messages provide a logical grouping of frames. For example, an HTTP request or response is represented as a message,
which consists of multiple frames (e.g., headers, data, etc.).

## 服务熔断

* 服务熔断是一种软件设计模式，用于分布式系统中处理服务调用失败的情况，可以防止被调用服务因为频繁失败被压垮。它借鉴了电路中的断路器原理，通过监控服务调用的失败率等条件来决定是否阻止进一步的调用，以保护系统免受过载。
阻止上游系统连续不断的调用 
* 服务熔断器有三个主要状态：关闭（Closed）、半开（HalfOpen）和打开（Open），分别对应不同的保护策略。当服务调用失败次数超过阈值时，熔断器打开，阻止服务调用。在一定时间后，熔断器尝试半开状态，允许少量请求通过以测试服务恢复情况。如果服务恢复，熔断器关闭；如果失败，熔断器保持打开状态。
* 在 go 语言里可以使用 sentinel-golang 库实现熔断功能。

## 动态扩缩容节点