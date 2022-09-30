/*
 * @lc app=leetcode.cn id=560 lang=cpp
 *
 * [560] 和为 K 的子数组
 */
#include <unordered_map>
#include <vector>

using namespace std;
// @lc code=start
class Solution {
public:
  int subarraySum(vector<int> &nums, int k) {
    unordered_map<int, int> m;
    int ans = 0;
    int pre = 0;
    m[0] = 1;
    for (int n : nums) {
      pre += n;
      if (m.find(pre - k) != m.end()) {
        ans += m[pre - k];
      }
      m[pre]++;
    }
    return ans;
  }
};
// @lc code=end
