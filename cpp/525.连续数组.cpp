/*
 * @lc app=leetcode.cn id=525 lang=cpp
 *
 * [525] 连续数组
 */
#include <unordered_map>
#include <vector>
using namespace std;
// @lc code=start
class Solution {
public:
  // n = [0, 1, 0]
  // i : 0, 1, 2
  // p : 0, 1, 1
  // 2(p1 - p2)= i1 - i2 ==> p1 - i1/2 = p2 - i2/2 + 1/2
  int findMaxLength(vector<int> &nums) {
    unordered_map<int, int> m;
    int ans = 0;
    int pre = 0;
    m[0] = -1;

    for (int i = 0; i < nums.size(); i++) {
      int n = nums[i];
      if (n == 0) {
        n = -1;
      }
      pre += n;
      if (m.find(pre) != m.end()) {
        int l = i - m[pre];
        if (ans < l) {
        ans = l;
        }
      } else {
      m[pre] = i;
      }
      

    }

    return ans;
  }
};
// @lc code=end
