/*
 * @lc app=leetcode.cn id=724 lang=cpp
 *
 * [724] 寻找数组的中心下标
 */
#include <vector>
using namespace std;
// @lc code=start
class Solution {
public:
    int pivotIndex(vector<int>& nums) {
        int sum = 0;
        for (int n : nums) {
        sum += n;
        }

        int pre = 0;

        for (int i = 0; i <  nums.size(); i++) {
        if (sum - nums[i] == pre * 2) {
        return i;
        }
        pre += nums[i];
        }

        return -1;
    }
};
// @lc code=end

