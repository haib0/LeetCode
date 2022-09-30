/*
 * @lc app=leetcode.cn id=354 lang=cpp
 *
 * [354] 俄罗斯套娃信封问题
 */
#include <algorithm>
#include <vector>
using namespace std;
// @lc code=start
class Solution {
public:
  int maxEnvelopes(vector<vector<int>> &envelopes) {
    if (envelopes.empty()) {
      return 0;
    }
    
    sort(envelopes.begin(), envelopes.end(),
         [](vector<int> e1, vector<int> e2) {
           return e1[0] < e2[0] || (e1[0] == e2[0] && e1[1] > e2[1]);
         });

    int ans = 0;
    vector<int> piles(envelopes.size());
    for (int i = 0; i < envelopes.size(); ++i) {
      int n = envelopes[i][1];

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
        ans++;
      }
      piles[left] = n;
    }
    return ans;
  }
};
// @lc code=end
