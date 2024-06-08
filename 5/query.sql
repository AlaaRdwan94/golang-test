WITH RankedEmployees AS (
    SELECT 
        e.Name AS Employee,
        e.Salary,
        d.Name AS Department,
        RANK() OVER (PARTITION BY e.departmentId ORDER BY e.Salary DESC) AS Rank
    FROM 
        Employee e
    JOIN 
        Department d ON e.departmentId = d.id
)
SELECT 
    Department,
    Employee,
    Salary
FROM 
    RankedEmployees
WHERE 
    Rank <= 3
ORDER BY 
    Department, Salary DESC;
