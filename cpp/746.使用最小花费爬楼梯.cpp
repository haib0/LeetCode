/*
 * @lc app=leetcode.cn id=746 lang=cpp
 *
 * [746] 使用最小花费爬楼梯
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
public:
  int minCostClimbingStairs(vector<int> &cost) {
    int n = cost.size();
    if (n == 1) {
      return 0;
    }
    vector<int> dp = vector<int>(n + 1);
    dp[0] = 0;
    dp[1] = cost[0];
    for (int i = 2; i <= n; i++) {
      auto x = min(dp[i - 1], dp[i - 2]);
      dp[i] = x + cost[i - 1];
    }
    return min(dp[n], dp[n - 1]);
  }
};
// @lc code=end
