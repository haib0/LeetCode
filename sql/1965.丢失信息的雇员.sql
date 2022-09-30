--
-- @lc app=leetcode.cn id=1965 lang=mysql
--
-- [1965] 丢失信息的雇员
--

-- @lc code=start
# Write your MySQL query statement below

select employee_id
from
    (select * from Employees left join Salaries using(employee_id)
    union
    select * from Employees right join Salaries using(employee_id)
    ) as T
where T.name is null or T.salary is null
order by employee_id;

-- @lc code=end

