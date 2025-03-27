/*
Tabel: logs
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| num         | varchar |
+-------------+---------+
id is the primary key for this table.
id is an autoincrement column.

Write an SQL query to find all numbers that appear at least three times consecutively.

Return the result table in any order.

The query result format is in the following example.
*/

SELECT DISTINCT l1.num AS ConsecutiveNums
FROM `logs` l1,
     `logs` l2,
     `logs` l3
WHERE l1.id + 1 = l2.id
  AND l2.id + 1 = l3.id
  AND l1.num = l2.num
  AND l2.num = l3.num;

-- 展示的结果根据最后的order排列  partition by Num order by Id
SELECT Id,
       Num,
       row_number() over(order by id) - row_number() over(partition by Num order by Id) as SerialNumberSubGroup
FROM logs

-- 展示的结果根据最后的order排列  order by id
SELECT Id,
       Num,
       row_number() over(partition by Num order by Id) + row_number() over(order by id)  as SerialNumberSubGroup
FROM logs

-- 展示的结果根据最后的order排列  order by id
SELECT Id,
       Num,
       row_number() over(partition by Num order by Id) as part,
       row_number() over(order by id) as SerialNumberSubGroup
FROM logs

-- 展示的结果根据最后的order排列  partition by Num order by Id
SELECT Id,
       Num,
        row_number() over(order by id) as SerialNumberSubGroup,
       row_number() over(partition by Num order by Id) as part
FROM logs