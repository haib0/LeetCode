--
-- @lc app=leetcode.cn id=1484 lang=mysql
--
-- [1484] 按日期分组销售产品
--

-- @lc code=start
# Write your MySQL query statement below

select sell_date, count(distinct product) as num_sold, group_concat(distinct product) as products
from Activities
group by sell_date;

-- @lc code=end

