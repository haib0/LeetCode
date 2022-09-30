--
-- @lc app=leetcode.cn id=1757 lang=mysql
--
-- [1757] 可回收且低脂的产品
--
-- @lc code=start
# Write your MySQL query statement below
SELECT product_id
FROM Products
WHERE low_fats = 'Y'
    AND recyclable = 'Y';
-- @lc code=end