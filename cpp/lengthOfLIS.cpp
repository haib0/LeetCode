// https://labuladong.github.io/algo/3/24/75/

#include <vector>
using namespace std;

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

int lengthOfLIS_dp(vector<int> &nums) {
  int ans = 1;
  vector<int> dp(nums.size());

  for (int i = 0; i < nums.size(); ++i) {
    int dpi = 1;
    for (int j = i - 1; j >= 0; --j) {
      if (nums[j] < nums[i] && dpi < dp[j] + 1) {
        dpi = dp[j] + 1;
      }
    }
    if (ans < dpi) {
      ans = dpi;
    }
    dp[i] = dpi;
  }

  return ans;
}