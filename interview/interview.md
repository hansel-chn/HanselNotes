1. 说广度和深度，
   实际项目，不规范，实际部署方面，资源分配

公司为初创公司，偏向于业务实现，流程上的一些东西并不标准，
、很多知识是自学得来的，所以更偏向于理论。（在以往实际落地的）

1.服务熔断，服务降级,服务限流

2. 服务注册发现

kafka消息丢失之类设置详细

##             

认识更加厉害的人，处理很多事情

如果您对我的项目的一些细节更为关注，请接下来对我进行更细致的提问。 或者看您需不需要我整体上的讲一下我的项目

docker的原理

## 高并发思路

日活跃用户数，二八法则估算qps，实际上线注意峰值流量

## 一个正常中等公司内部网络应用，2000人左右，内部流量，qps多少，需要几台nginx

对于一个中等规模的公司内部网络应用，2000人左右的用户，QPS（每秒查询数）和所需的 Nginx 实例数量可以通过以下方式估算：

### 1. **QPS 估算**

- **活跃用户数**：假设 2000 人中有 20% 是同时在线的活跃用户（即 400 人）。
- **每秒请求数**：假设每个活跃用户每分钟发起 10 次请求，则每秒的请求数为：
  \[
  QPS = \frac{400 \times 10}{60} \approx 67
  \]
- **峰值流量**：考虑到流量的波动，通常需要预留 2-3 倍的峰值冗余。假设峰值流量为 3 倍，则峰值 QPS 为：
  \[
  QPS_{peak} = 67 \times 3 = 201
  \]

### 2. **Nginx 性能**

- Nginx 在处理静态资源时，单台服务器的 QPS 可以轻松达到 10,000+（取决于硬件配置）。
- 如果涉及动态内容（如代理后端服务），单台 Nginx 的性能通常在 2,000-5,000 QPS 左右。

### 3. **所需 Nginx 数量**

- 假设单台 Nginx 的性能为 2,000 QPS，且需要考虑高可用性（HA），至少需要 2 台 Nginx 服务器（1 主 1 备）。
- 如果需要更高的冗余，可以部署 3 台或更多 Nginx 服务器。

### 4. **总结**

- **估算 QPS**：约 67（峰值 201）。
- **推荐 Nginx 数量**：2-3 台（1 主 1 备，或 1 主 2 备）。
- **硬件配置建议**：
    - CPU：4 核或以上。
    - 内存：4 GB 或以上。
    - 网络：千兆网卡。

此外，建议对 Nginx 进行性能调优（如增加 `worker_connections` 和 `worker_processes`），并使用负载均衡（如 LVS 或
Keepalived）来提高可靠性。

## lvs需要多台吗，还是一台就够，还是一台就完全可以处理

The number of LVS (Linux Virtual Server) instances required depends on your system's traffic, redundancy needs, and
scalability requirements:

1. **Single LVS Instance**:
    - A single LVS instance can handle significant traffic if properly configured and running on capable hardware.
    - However, it introduces a **single point of failure**. If the LVS instance goes down, the entire system becomes
      unavailable.

2. **Multiple LVS Instances**:
    - For **high availability (HA)**, it's recommended to deploy at least two LVS instances in an **active-passive** or
      **active-active** configuration.
    - Tools like **Keepalived** or **Heartbeat** are commonly used to manage failover between LVS instances.
    - This setup ensures that if one LVS instance fails, the other can take over seamlessly.

3. **Scalability**:
    - If traffic increases beyond the capacity of a single LVS instance, you can scale horizontally by adding more LVS
      instances and using DNS load balancing or anycast to distribute traffic among them.

### Recommendation:

- For small to medium traffic, **one LVS instance with a backup** (using Keepalived) is sufficient.
- For high traffic or critical systems, deploy **multiple LVS instances** with redundancy and load balancing.

## 单个lvs一般能应付多大流量，qps

单个 LVS 的流量和 QPS 能力主要取决于以下几个因素：

### 1. **硬件性能**

- **CPU**: LVS 运行在内核态，依赖 CPU 的处理能力。现代多核 CPU 可以显著提高 LVS 的并发处理能力。
- **内存**: LVS 本身对内存需求较低，但需要足够的内存来处理连接跟踪表（conntrack）。
- **网络带宽**: 受限于服务器的网卡性能，例如：
    - 1 Gbps 网卡理论上可以处理约 125 MB/s 的流量。
    - 10 Gbps 网卡理论上可以处理约 1.25 GB/s 的流量。

### 2. **LVS 模式**

- **NAT 模式**: LVS 需要对每个数据包进行地址转换，CPU 和内存开销较大，适合中小规模流量。
- **DR 模式**: LVS 只处理入站流量，出站流量由后端服务器直接返回客户端，性能更高，适合大规模流量。
- **TUN 模式**: 类似 DR 模式，但通过隧道转发，适合跨网络的高性能场景。

### 3. **QPS 理论值**

- **NAT 模式**: 单台 LVS 通常可以处理 100,000 - 300,000 QPS，具体取决于硬件和内核优化。
- **DR/TUN 模式**: 单台 LVS 可以处理 500,000 - 1,000,000 QPS 或更高，瓶颈通常是网卡和网络带宽。

### 4. **优化建议**

- **内核参数调优**: 增加连接跟踪表大小（`net.netfilter.nf_conntrack_max`）、优化 TCP 参数（如 `tcp_tw_reuse`）。
- **网卡优化**: 使用多队列网卡（如 SR-IOV）和中断绑定（IRQ Affinity）。
- **负载均衡算法**: 根据业务场景选择合适的调度算法（如 `rr`、`wlc`）。
- **高可用性**: 使用 Keepalived 配置主备 LVS，避免单点故障。

### 总结

单台 LVS 在高性能硬件和优化配置下，理论上可以处理数百万 QPS 和接近网卡带宽的流量（如 10 Gbps）。实际性能需根据具体业务场景和硬件条件测试验证。

## 怎么考虑nginx本身的性能和后台服务的性能，怎么结合计算网关的性能，qps

To evaluate and combine the performance of Nginx and backend services to calculate the overall gateway performance (
QPS), follow these steps:

---

### 1. **Evaluate Nginx Performance**

- **Static Content**: If Nginx serves static files, its QPS is primarily limited by network bandwidth and hardware (
  e.g., CPU, memory).
- **Proxy Mode**: When Nginx proxies requests to backend services, its QPS depends on:
    - Connection handling (`worker_connections`, `worker_processes`).
    - TLS overhead (if enabled).
    - Logging and buffering configurations.

**Key Metrics**:

- Measure Nginx's maximum QPS using tools like `wrk` or `ab` under different configurations.
- Optimize Nginx settings (e.g., `keepalive_timeout`, `worker_connections`).

---

### 2. **Evaluate Backend Service Performance**

- Measure the backend service's QPS independently by simulating requests directly to the service.
- Consider:
    - Database query performance.
    - Business logic execution time.
    - Latency and throughput.

**Key Metrics**:

- Average response time per request.
- Maximum concurrent requests the backend can handle.

---

### 3. **Combine Nginx and Backend Performance**

- Nginx acts as a gateway, so its QPS is limited by the slower component (Nginx or backend).
- Use the following formula to estimate the combined QPS:

\[
QPS_{gateway} = \min(QPS_{nginx}, QPS_{backend})
\]

- If Nginx is the bottleneck, optimize its configuration.
- If the backend is the bottleneck, scale backend services horizontally or optimize their performance.

---

### 4. **Consider Network and Resource Overheads**

- Ensure sufficient network bandwidth between Nginx and backend services.
- Monitor CPU, memory, and disk I/O usage on both Nginx and backend servers.

---

### 5. **Test and Validate**

- Use load testing tools (e.g., `wrk`, `JMeter`) to simulate real-world traffic.
- Gradually increase the load to identify bottlenecks.

---

### Example: Combined QPS Calculation

- Nginx can handle 50,000 QPS (proxy mode).
- Backend services can handle 10,000 QPS.
- Gateway QPS = `min(50,000, 10,000) = 10,000`.

To improve, scale backend services or optimize their performance.

## 两条记录A B死锁了，怎么处理

package main

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
)

func main() {
db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
if err != nil {
panic(err)
}
defer db.Close()

       for i := 0; i < 3; i++ { // 重试最多3次
           tx, err := db.Begin()
           if err != nil {
               panic(err)
           }

           _, err = tx.Exec("UPDATE table_name SET column = value WHERE id = ?", 1)
           if err != nil {
               tx.Rollback()
               if isDeadlockError(err) {
                   fmt.Println("Deadlock detected, retrying...")
                   continue
               }
               panic(err)
           }

           _, err = tx.Exec("UPDATE table_name SET column = value WHERE id = ?", 2)
           if err != nil {
               tx.Rollback()
               if isDeadlockError(err) {
                   fmt.Println("Deadlock detected, retrying...")
                   continue
               }
               panic(err)
           }

           err = tx.Commit()
           if err != nil {
               panic(err)
           }
           fmt.Println("Transaction committed successfully")
           break
       }

}

func isDeadlockError(err error) bool {
return err != nil && (err.Error() == "Error 1213: Deadlock found when trying to get lock")
}

## 死锁解决问题

1. 资源是否占用过多

mysql

* 数据库会自动检测死锁错误并重试，
* 锁的的顺序是否可以优化保证避免循环依赖
* 分布式锁协调资源访问

## 零信任

永不信任，始终验证