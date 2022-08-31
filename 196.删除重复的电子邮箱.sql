--
-- @lc app=leetcode.cn id=196 lang=mysql
--
-- [196] 删除重复的电子邮箱
--

-- @lc code=start
# Please write a DELETE statement and DO NOT write a SELECT statement.
# Write your MySQL query statement below

delete P1
from Person P1, Person P2
where P1.email = P2.email and P1.id > P2.id

-- @lc code=end

