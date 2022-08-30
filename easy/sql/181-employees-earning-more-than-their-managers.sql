# Write your MySQL query statement below
SELECT
    emp.name "Employee"
FROM
    employee emp
        JOIN
    employee mng
    ON
            emp.managerId = mng.id
WHERE
        emp.salary > mng.salary