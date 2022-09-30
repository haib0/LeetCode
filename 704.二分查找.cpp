/*
 * @lc app=leetcode.cn id=704 lang=cpp
 *
 * [704] 二分查找
 */
#include <vector>

using namespace std;
// @lc code=start
class Solution {
public:
  int search(vector<int> &nums, int target) {
    int head = 0, tail = nums.size() - 1;
    while (head <= tail) {
      int mid = (tail + head) / 2;
      if (nums[mid] < target) {
        head = mid + 1;
        continue;
      }
      if (nums[mid] > target) {
        tail = mid - 1;
        continue;
      }

      return mid;
    }

    return -1;
  }
};
// @lc code=end
