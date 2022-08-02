/*
 * @lc app=leetcode.cn id=1991 lang=cpp
 *
 * [1991] 找到数组的中间位置
 */
#include <vector>
using namespace std;
// @lc code=start
class Solution {
public:
  int findMiddleIndex(vector<int> &nums) {
    int sum = 0;
    for (int n : nums) {
      sum += n;
    }

    int pre = 0;

    for (int i = 0; i < nums.size(); i++) {
      if (sum - nums[i] == pre * 2) {
        return i;
      }
      pre += nums[i];
    }

    return -1;
  }
};
// @lc code=end
