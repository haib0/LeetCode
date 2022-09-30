--
-- @lc app=leetcode.cn id=1795 lang=mysql
--
-- [1795] 每个产品在不同商店的价格
--

-- @lc code=start
# Write your MySQL query statement below

select product_id, 'store1' store, store1 price
from Products
where store1 is not null

union

select product_id, 'store2' store, store2 price
from Products
where store2 is not null

union

select product_id, 'store3' store, store3 price
from Products
where store3 is not null

order by product_id;

-- @lc code=end

