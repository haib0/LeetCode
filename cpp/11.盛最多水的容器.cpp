/*
 * @lc app=leetcode.cn id=11 lang=cpp
 *
 * [11] 盛最多水的容器
 */
#include <vector>
using namespace std;
// @lc code=start
class Solution
{
public:
    int maxArea(vector<int> &height)
    {
        int maxArea = 0;

        int left = 0;
        int right = height.size() - 1;

        int area;
        while (left < right)
        {
            area = (right - left) * (height[left] < height[right] ? height[left] : height[right]);
            maxArea = area > maxArea ? area : maxArea;

            if (height[left] < height[right])
            {
                do
                {
                    left++;
                } while (left < right && height[left] <= height[left - 1]);
            }
            else
            {
                do
                {
                    right--;
                } while (left < right && height[right] <= height[right + 1]);
            }
        }

        return maxArea;
    }
};
// @lc code=end
