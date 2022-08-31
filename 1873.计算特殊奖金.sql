--
-- @lc app=leetcode.cn id=1873 lang=mysql
--
-- [1873] 计算特殊奖金
--

-- @lc code=start
# Write your MySQL query statement below

select employee_id, 
if(employee_id%2 != 0 and name not like 'M%', salary, 0) as bonus
from Employees
order by employee_id;

-- @lc code=end

