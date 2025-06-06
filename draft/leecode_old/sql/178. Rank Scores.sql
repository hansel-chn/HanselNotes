/*
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| score       | decimal |
+-------------+---------+
id is the primary key for this table.
Each row of this table contains the score of a game. Score is a floating point value with two decimal places.

Write an SQL query to rank the scores. The ranking should be calculated according to the following rules:

The scores should be ranked from the highest to the lowest.
If there is a tie between two scores, both should have the same ranking.
After a tie, the next ranking number should be the next consecutive integer value. In other words, there should be no holes between ranks.
Return the result table ordered by score in descending order.

The query result format is in the following example.
*/

select s.score as `score`, (select count(distinct p.score) from `Scores` as p where p.score >= s.score) as `rank`
from `Scores` as `s`
order by s.score desc

SELECT s.score AS `score`, dense_rank() over (ORDER BY s.score DESC ) AS `rank` FROM `Scores` AS `s`

