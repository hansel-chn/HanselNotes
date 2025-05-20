## 怎么绕过了age的间隙锁，不是已经锁定大于20的记录了？

绕过 `age` 的间隙锁（Gap Lock）通常是因为以下原因之一：

### 1. **插入操作使用了不同的索引**

InnoDB 的锁是基于索引的。如果插入操作使用了唯一索引（如主键 `id`），而间隙锁是基于非唯一索引（如 `age`），那么间隙锁不会阻止插入操作。例如：

- 查询 `SELECT * FROM users WHERE age > 20 FOR UPDATE` 会在 `age` 索引上加间隙锁。
- 插入 `INSERT INTO users (id, name, age) VALUES (101, 'Alice', 25)` 使用的是主键 `id` 索引，因此不会被 `age` 索引的间隙锁阻止。

### 2. **事务隔离级别不是 REPEATABLE READ**

间隙锁只在 **REPEATABLE READ** 隔离级别下生效。如果事务隔离级别是 **READ COMMITTED**，InnoDB
不会使用间隙锁，而是只锁定具体的行。这种情况下，插入操作不会被阻止。

### 3. **查询范围未覆盖插入的记录**

如果查询范围未完全覆盖插入的记录，间隙锁可能未锁定插入位置。例如：

```sql
SELECT *
FROM users
WHERE age > 20 FOR UPDATE;
```

- 如果表中没有 `age > 20` 的记录，InnoDB 可能只锁定了 `age` 索引的末尾间隙。
- 插入 `age = 25` 的记录可能未被锁定的间隙覆盖，因此插入成功。

### 4. **插入操作是由当前事务执行**

如果插入操作和查询操作在同一个事务中，InnoDB 不会阻止自己事务的插入操作。

---

### 如何避免绕过间隙锁

1. **确保查询范围覆盖所有可能的插入位置**：
    - 如果需要锁定所有 `age > 20` 的间隙，可以确保表中有一条记录满足条件，或者使用更明确的范围查询。

2. **使用主键或唯一索引加锁**：
    - 在查询时同时锁定主键或唯一索引，避免插入操作绕过非唯一索引的间隙锁。

3. **验证事务隔离级别**：
    - 确保事务隔离级别为 **REPEATABLE READ**，以启用间隙锁。

4. **明确锁定范围**：
    - 使用 `SELECT ... FOR UPDATE` 或 `LOCK IN SHARE MODE` 明确锁定需要保护的范围。

---

### 示例

确保间隙锁生效的查询：

```sql
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;
START TRANSACTION;
SELECT *
FROM users
WHERE age > 20 FOR UPDATE;
-- 确保锁定所有 age > 20 的间隙
```

## If another transaction inserts a new row (e.g., age = 25) and commits, the new row is not part of the original snapshot. However, if the first transaction re-executes the query without locking, it may see the new row, causing a phantom read. 为什么，是因为范围查询和单挑查询逻辑不一样？没有用到mvcc机制？为什么可见

The reason why the new row becomes visible in a **range query** but not in a **single-row query** lies in the way *
*MVCC (Multi-Version Concurrency Control)** and **locking mechanisms** handle these scenarios.

### Key Points:

1. **MVCC Mechanism**:
    - MVCC ensures that each transaction sees a consistent snapshot of the database as of the transaction's start time.
    - For **single-row queries**, MVCC uses the row's version information (transaction IDs) to determine visibility. If
      a row was inserted or modified after the transaction's snapshot was taken, it will not be visible to the current
      transaction.
    - For **range queries**, MVCC also ensures consistency for rows that existed at the time of the snapshot. However, *
      *new rows inserted into the range after the snapshot** are not inherently excluded unless additional locking is
      applied.

2. **Range Queries and Phantom Reads**:
    - A **range query** (e.g., `SELECT * FROM users WHERE age > 20`) does not lock the "gaps" between rows unless
      explicitly requested (e.g., using `FOR UPDATE` or `LOCK IN SHARE MODE`).
    - If another transaction inserts a new row (e.g., `age = 25`) and commits, this row is not part of the original
      snapshot. However, when the first transaction re-executes the range query, it evaluates the condition (`age > 20`)
      again, and the new row matches the condition, making it visible.

3. **Why Single-Row Queries Are Different**:
    - A **single-row query** (e.g., `SELECT * FROM users WHERE id = 1`) directly targets a specific row. MVCC ensures
      that only the version of the row visible at the time of the snapshot is returned.
    - Since no new rows can match the query condition (as the row's identity is fixed), phantom reads do not occur in
      single-row queries.

4. **Gap Locking in Range Queries**:
    - To prevent phantom reads in range queries, **gap locks** are required. These locks explicitly block the insertion
      of new rows into the range.
    - Without gap locks, the range query does not prevent other transactions from inserting rows that match the query
      condition.

### Example:

```sql
-- Transaction 1
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;
START TRANSACTION;
SELECT *
FROM users
WHERE age > 20;
-- Reads rows with age = 25, 30

-- Transaction 2
START TRANSACTION;
INSERT INTO users (id, age)
VALUES (3, 28); -- Inserts a new row
COMMIT;

-- Transaction 1
SELECT *
FROM users
WHERE age > 20; -- Now sees rows with age = 25, 28, 30 (phantom read)
COMMIT;
```

### Why the New Row is Visible:

- The **range query** evaluates the condition (`age > 20`) dynamically each time it is executed.
- Without gap locks, the new row (`age = 28`) is not excluded by MVCC because it matches the query condition and is
  committed.

To avoid this, you must explicitly lock the range using `FOR UPDATE` or use the `SERIALIZABLE` isolation level.