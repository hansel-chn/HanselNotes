-- Table: Employee
--
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | id           | int     |
-- | name         | varchar |
-- | salary       | int     |
-- | departmentId | int     |
-- +--------------+---------+
-- id is the primary key column for this table.
-- departmentId is a foreign key of the ID from the Department table.
-- Each row of this table indicates the ID, name, and salary of an employee. It also contains the ID of their department.
--
--
-- Table: Department
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | name        | varchar |
-- +-------------+---------+
-- id is the primary key column for this table.
-- Each row of this table indicates the ID of a department and its name.
--
--
-- Write an SQL query to find employees who have the highest salary in each of the departments.
--
-- Return the result table in any order.
--
-- The query result format is in the following example.

-- 内连接查询
SELECT d.name AS Department, e1.name AS Employee, e1.salary AS Salary
FROM `Employee` AS `e1`
         INNER JOIN (SELECT MAX(`salary`) AS `maxSalary`, e2.departmentId AS `departmentId`
                     FROM `Employee` AS `e2`
                     GROUP BY e2.departmentId) AS `t1` ON e1.salary = t1.maxSalary AND t1.departmentId = e1.departmentId
         INNER JOIN Department AS `d` ON d.id = e1.departmentId;


-- where in 匹配多个字段
SELECT Department.name AS 'Department', Employee.name AS 'Employee', Salary
FROM Employee
         JOIN
     Department ON Employee.DepartmentId = Department.Id
WHERE (Employee.DepartmentId, Salary) IN
      (SELECT DepartmentId,
              MAX(Salary)
       FROM Employee
       GROUP BY DepartmentId
      )
;
