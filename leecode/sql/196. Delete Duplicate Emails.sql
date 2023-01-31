-- Table: Person
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | email       | varchar |
-- +-------------+---------+
-- id is the primary key column for this table.
-- Each row of this table contains an email. The emails will not contain uppercase letters.
--
--
-- Write an SQL query to delete all the duplicate emails, keeping only one unique email with the smallest id. Note that you are supposed to write a DELETE statement and not a SELECT one.
--
-- After running your script, the answer shown is the Person table. The driver will first compile and run your piece of code and then show the Person table. The final order of the Person table does not matter.
--
-- The query result format is in the following example.

DELETE
`p1`
FROM `Person` AS `p1` INNER JOIN `Person` AS `p2`
ON p1.email = p2.email and p1.id > p2.id

-- 这样不可以，子查询的结果不能用来删除数据，如下错误
DELETE
FROM `Person` AS `p2`
WHERE p2.id not exists (SELECT min(p.id) FROM `Person` AS `p` GROUP BY p.email having COUNT(1) > 1)

-- 需要把子查询的结果用临时表保存  而且不能用exist
DELETE
`p2`
FROM `Person` AS `p2`
WHERE p2.id not in
(SELECT *
FROM (SELECT min(p.id) FROM `Person` AS `p` GROUP BY p.email) AS `t3`)

-- 不能在FROM子句中指定更新的目标表'p2'
DELETE `p2`
FROM `Person` AS `p2`
WHERE not exists(SELECT 1 FROM `Person` AS `p` GROUP BY p.email having min(p.id) = p2.id)

-- exists 原理
-- exists做为where条件时,是先对where 前的主查询询进行查询,然后用主查询的结果一个一个的代入exists的查询进行判断，如果为真则输出当前这一条主查询的结果，否则不输出。
-- 当子查询结果非常大时，EXISTS子句比IN快得多。相反，当子查询结果非常小时，IN子句比EXISTS快。
-- 如果使用IN操作符，SQL引擎将扫描从内部查询中获取的所有记录。另一方面，如果我们使用EXISTS, SQL引擎将在找到匹配项后立即停止扫描过程。
-- https://blog.csdn.net/a_hui_tai_lang/article/details/81146635?spm=1001.2101.3001.6650.5&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-5-81146635-blog-108504594.pc_relevant_3mothn_strategy_and_data_recovery&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-5-81146635-blog-108504594.pc_relevant_3mothn_strategy_and_data_recovery&utm_relevant_index=10
-- 实际上优化器可能做了优化