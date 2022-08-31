--
-- @lc app=leetcode.cn id=1527 lang=mysql
--
-- [1527] 患某种疾病的患者
--

-- @lc code=start
# Write your MySQL query statement below

select *
from Patients
where conditions regexp '\\bDIAB1';
-- where conditions like "DIAB1%" or conditions like "% DIAB1%";

-- @lc code=end

