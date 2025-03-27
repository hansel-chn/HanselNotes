-- Table: Employee
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | name        | varchar |
-- | salary      | int     |
-- | managerId   | int     |
-- +-------------+---------+
-- id is the primary key column for this table.
-- Each row of this table indicates the ID of an employee, their name, salary, and the ID of their manager.
--
--
-- Write an SQL query to find the employees who earn more than their managers.
--
-- Return the result table in any order.
--
-- The query result format is in the following example.

SELECT e.name AS `Employee` FROM `EMPLOYEE` AS `e` WHERE e.salary > (SELECT e1.salary FROM `EMPLOYEE` AS `e1` WHERE  e1.id = e.managerId)

SELECT e1.name AS `Employee` FROM `EMPLOYEE` AS `e1` INNER JOIN `EMPLOYEE` AS `e2` ON  e1.managerId = e2.id AND e1.salary > e2.salary;