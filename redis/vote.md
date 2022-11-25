#主从复制

#哨兵机制选举
哨兵Sentinel监控，当master宕机，进行选举自动切换（足量Sentinel确认主服务器下线，包括主从的Sentinel）
实现读写分离，但是未分布式存储，浪费内存

#集群模式and选举
插槽实现分布式存储
多个master，多个slave
半数以上主节点ping失败认为目标主节点宕机，切换节点（足量主节点确认主服务器下线）类似哨兵原理
（选举从节点也是足量主节点确认）<mark>集群是master node来判断下线，切换。Sentinel机制是主从的Sentinel判断</mark>