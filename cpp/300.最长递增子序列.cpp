/*
 * @lc app=leetcode.cn id=300 lang=cpp
 *
 * [300] 最长递增子序列
 */
#include <vector>

using namespace std;
// @lc code=start
class Solution {
public:
  int lengthOfLIS(vector<int> &nums) {
    int ans = 0;
    vector<int> piles(nums.size());

    for (int i = 0; i < nums.size(); ++i) {
      int n = nums[i];

      // binary search n, left bound
      int left = 0, right = ans;
      while (left < right) {
        int mid = (left + right) / 2;
        if (piles[mid] < n) {
          left = mid + 1;
          continue;
        }
        if (piles[mid] > n) {
          right = mid;
          continue;
        }
        right = mid;
      }

      if (left == ans) {
        ++ans;
      }
      piles[left] = n;
    }
    return ans;
  }
};
// @lc code=end
