/*
 * @lc app=leetcode.cn id=69 lang=cpp
 *
 * [69] Sqrt(x)
 */

// @lc code=start
class Solution
{
public:
    int mySqrt(int x)
    {
        int ans = 0;
        while (true)
        {

            if ((long long)ans * ans > x)
            {
                return ans - 1;
            }
            ans++;
        }

        return ans;
    }
};
// @lc code=end
