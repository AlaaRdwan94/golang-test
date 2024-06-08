Write SQL Queries for Department of Top Three Salaries

Explanation:

We first use a common table expression (CTE) named RankedEmployees to select employee names, salaries, department names, and rank of salaries within each department. The RANK() function is used to assign a rank to each employee's salary within their respective department, with the highest salary receiving a rank of 1.
We then select from the RankedEmployees CTE where the rank is less than or equal to 3, indicating the top three earners in each department.
Finally, we order the result set by department name and salary in descending order.
This query will return a result set containing the department name, employee name, and salary of the top three earners in each department.

