--
-- @lc app=leetcode.cn id=608 lang=mysql
--
-- [608] 树节点
--

-- @lc code=start
# Write your MySQL query statement below

select id,
    case 
        when p_id is null then 'Root'
        when id in (select p_id from Tree) then 'Inner'
        else 'Leaf'
    end as type
from Tree;

-- @lc code=end

