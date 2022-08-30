--
-- @lc app=leetcode.cn id=584 lang=mysql
--
-- [584] 寻找用户推荐人
--
-- @lc code=start
# Write your MySQL query statement below
SELECT name
FROM customer
WHERE referee_id is null || referee_id != 2;
-- @lc code=end