# Redis

# redis类型

### string

内部存储形式,SDS简单动态字符串
https://redis.io/docs/latest/operate/oss_and_stack/reference/internals/internals-sds/

设计SDS的目的，

1. 减少内存分配次数，
2. 利用CPU缓存提升性能

何时可用:

1. 大小在一定限制范围内，
2. 数据为只读不做改变

## distributed lock

### consul

### zookeeper

### Redlock

1. 确保高可用和错误容忍度
2. 通过建立一种机制，解决如何在多个redis实例上获取大部分实例上所得问题，确保了同一个lock在不同实例上被一致性的识别

处理了计算了，获取锁的延迟时间，设计了合理的获取锁的超时时间，以获取大部分实例上的锁从而获得分布式锁（半数以上）
核心在于大部分instance对锁的状态达成一致

算法上从数学上好像证明了严谨性，但是更推荐使用raft等协议等共识算法

## CAP理论

一致性：consistency
高可用：available
分区容错性：partition tolerance
由于分布式系统天然存在分区容错性，所以需要在一致性和高可用中做出权衡

etcd相关内容：怎么“同时”实现两个

## ziplist

连锁更新问题
ziplist 设计出的紧凑型数据块可以有效利用内存，但在更新上，由于每一个 entry 都保留了前一个 entry 的 prevlen
长度，因此在插入或者更新时可能会出现连锁更新，这是一个影响效率的大问题。

因此，接着又设计出 「链表 + ziplist」组成的 quicklist 结构来避免单个 ziplist 过大，可以有效降低连锁更新的影响面。

但 quicklist 本质上不能完全避免连锁更新问题，因此，又设计出与 ziplist 完全不同的内存紧凑型结构 listpack，继续往下看～
在 listpack 中，因为每个列表项只记录自己的长度，而不会像 ziplist 中的列表项那样，会记录前一项的长度。所以，当在 listpack
中新增或修改元素时，实际上只会涉及每个列表项自己的操作，而不会影响后续列表项的长度变化，这就避免了连锁更新。

## 红黑树 跳表和b+树

在内存中，跳表优于 B+ 树的原因主要体现在以下几个方面：

### 1. **实现简单**

- 跳表的实现比 B+ 树简单，代码复杂度低，维护成本小。
- B+ 树需要复杂的节点分裂、合并逻辑，而跳表只需要维护多层索引。

### 2. **动态性更好**

- 跳表在插入和删除时的复杂度为 \(O(\log N)\)，且操作简单。
- B+ 树在插入和删除时可能需要频繁的节点分裂或合并，操作复杂度较高。

### 3. **内存访问模式**

- 跳表的节点是链表结构，内存分布较为分散，但在内存中随机访问的代价较低。
- B+ 树的节点是树形结构，可能会导致更多的内存页访问，尤其在节点分裂或合并时。

### 4. **范围查询效率**

- 跳表的多层索引结构使得范围查询高效，时间复杂度为 \(O(\log N + M)\)，其中 \(M\) 是结果集大小。
- B+ 树的范围查询也高效，但在内存中，跳表的链表结构更适合顺序扫描。

### 5. **内存占用**

- 跳表的节点结构简单，内存占用较低。
- B+ 树的节点需要存储更多的元数据（如子节点指针），内存开销较大。

### 6. **并发支持**

- 跳表天然支持分层并发控制，可以对不同层的索引进行独立加锁，提升并发性能。
- B+ 树的并发控制较为复杂，需要对节点进行精细化的锁管理。

### 总结

跳表在内存中优于 B+ 树的主要原因是其实现简单、动态性好、内存访问模式更适合内存环境，以及更低的内存占用和更好的并发支持。B+
树更适合磁盘存储场景，而跳表更适合内存中的高效操作。

红黑树相比普通平衡二叉树（如AVL树）的主要优点在于其旋转次数较少，以下是原因和优点的具体分析：

## 红黑树vs平衡二叉树

### 1. **旋转次数少的原因**

- **红黑树**：
    - 红黑树通过引入颜色（红色和黑色）来放宽平衡条件，允许树的高度在一定范围内不完全平衡。
    - 插入或删除节点时，最多需要 **O(log N)** 次旋转，通常旋转次数较少。
- **AVL树**：
    - AVL树要求严格的平衡条件（左右子树高度差不超过1），因此在插入或删除节点时，可能需要更多的旋转来维持平衡。

### 2. **红黑树的优点**

- **插入和删除效率更高**：
    - 由于红黑树的平衡条件较宽松，插入和删除操作需要的调整（旋转和重新着色）较少，性能更优。
    - AVL树在频繁插入和删除的场景中，旋转次数较多，性能可能受到影响。

- **适合写多读少的场景**：
    - 红黑树在插入和删除时的效率更高，适合写操作较多的场景。
    - AVL树在读操作性能稍优，但写操作开销较大。

- **实现复杂度较低**：
    - 红黑树的实现相对简单，代码复杂度低于AVL树，维护成本较低。

### 3. **适用场景的差异**

- **红黑树**：
    - 更适合写操作频繁的场景，如操作系统中的调度器、Java的`TreeMap`和`TreeSet`。
- **AVL树**：
    - 更适合读操作频繁的场景，因为其严格平衡性使得查找操作性能稍优。

### 总结

红黑树通过放宽平衡条件，减少了插入和删除时的旋转次数，提升了写操作的效率，同时实现复杂度较低，适合写多读少的场景。相比之下，AVL树的严格平衡性使其在读操作性能上稍优，但写操作开销较大。

## 并发访问的key数量很多，因为更新问题每个都加上锁，不会损耗很多空间吗

在高并发场景下，为每个 key 加锁确实可能会增加一定的内存开销，但这种开销通常是可控的。以下是一些优化策略，可以减少空间和性能损耗：

### 1. **使用分段锁（Shard Locking）**

- 将 key 分组，每组使用一个锁，而不是为每个 key 单独创建锁。
- 通过哈希函数将 key 映射到特定的锁，从而减少锁的数量。

示例代码：

```go
package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

const shardCount = 16

type ShardedLock struct {
	locks [shardCount]sync.Mutex
}

func (s *ShardedLock) getLock(key string) *sync.Mutex {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return &s.locks[hash.Sum32()%shardCount]
}

func main() {
	shardedLock := &ShardedLock{}

	key := "exampleKey"
	lock := shardedLock.getLock(key)

	lock.Lock()
	fmt.Println("Locked for key:", key)
	lock.Unlock()
	fmt.Println("Unlocked for key:", key)
}
```

### 2. **使用 Redis 分布式锁**

- 如果锁的粒度需要更细，可以使用 Redis 分布式锁（如 `SETNX` 或 `Redlock`），避免在本地内存中存储大量锁。

### 3. **锁的生命周期管理**

- 确保锁的生命周期尽可能短，避免长时间持有锁。
- 使用锁的过期时间（如 Redis 锁的 TTL）来防止死锁。

### 4. **批量操作**

- 如果多个 key 的更新可以合并，尝试批量操作，减少锁的数量和频率。

通过这些优化，可以有效减少锁的空间和性能开销，同时保证并发访问的正确性。

## golang mysql数据 定时更新到redis缓存需要注意什么

When periodically updating MySQL data into Redis cache in Go, you need to consider the following aspects:

### **Key Considerations**

1. **Concurrency**:

- Ensure thread-safe access to Redis during updates.
- Use synchronization mechanisms like `sync.Mutex` or atomic operations if necessary.

2. **Cache Consistency**:

- Avoid stale data by ensuring updates are atomic.
- Use a temporary Redis key or pipeline to update data and swap keys after completion.

3. **Error Handling**:

- Handle database connection errors and query failures gracefully.
- Implement retry logic for transient errors.

4. **Performance**:

- Optimize database queries to minimize load on MySQL.
- Use batch queries or pagination for large datasets.
- Use Redis pipelines for bulk updates to reduce network overhead.

5. **Cache Expiration**:

- Set appropriate TTL (Time-To-Live) for Redis keys to prevent stale data.

6. **Update Frequency**:

- Choose an appropriate update interval based on the application's requirements and database load.

7. **Monitoring**:

- Log update operations and monitor for failures or delays.

8. **Graceful Shutdown**:

- Ensure the update process can handle application shutdowns without leaving Redis in an inconsistent state.

### **Example Code**

Below is an example of periodically updating MySQL data into Redis cache:

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"context"
)

var ctx = context.Background()

func main() {
	// MySQL connection
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	// Start periodic cache update
	ticker := time.NewTicker(10 * time.Second) // Update every 10 seconds
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			updateCache(db, rdb)
		}
	}()

	// Simulate application running
	select {}
}

func updateCache(db *sql.DB, rdb *redis.Client) {
	// Fetch data from MySQL
	rows, err := db.Query("SELECT id, value FROM your_table")
	if err != nil {
		log.Printf("Failed to query MySQL: %v", err)
		return
	}
	defer rows.Close()

	// Use Redis pipeline for bulk updates
	pipe := rdb.Pipeline()
	for rows.Next() {
		var id int
		var value string
		if err := rows.Scan(&id, &value); err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		pipe.Set(ctx, fmt.Sprintf("cache_key:%d", id), value, 0) // No expiration
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Printf("Failed to update Redis cache: %v", err)
		return
	}

	fmt.Println("Cache updated successfully")
}
```

### **Explanation**

1. **Redis Pipeline**:

- Reduces network overhead by batching Redis commands.

2. **Atomic Updates**:

- Ensures consistency by updating Redis keys in bulk.

3. **Error Handling**:

- Logs errors during MySQL queries or Redis updates.

4. **Periodic Updates**:

- A `time.Ticker` triggers updates at regular intervals.

This approach ensures efficient and consistent cache updates while minimizing the impact on MySQL and Redis performance.