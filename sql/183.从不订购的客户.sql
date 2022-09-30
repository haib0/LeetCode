--
-- @lc app=leetcode.cn id=183 lang=mysql
--
-- [183] 从不订购的客户
--
-- @lc code=start
# Write your MySQL query statement below
-- METHOD 0
SELECT Name as Customers
FROM Customers
WHERE Id NOT IN (
        SELECT CustomerId
        FROM Orders
    );
-- METHOD 2
SELECT Name as Customers
FROM Customers
    LEFT JOIN Orders ON Customers.Id = Orders.CustomerId
WHERE Orders.CustomerId IS NULL;
-- @lc code=end