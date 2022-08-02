/*
 * @lc app=leetcode.cn id=35 lang=cpp
 *
 * [35] 搜索插入位置
 */
#include <vector>

using namespace std;
// @lc code=start
class Solution {
public:
  int searchInsert(vector<int> &nums, int target) {
    // left bound
    int head = 0, tail = nums.size();
    while (head < tail) {
      int mid = (tail + head) / 2;

      if (nums[mid] < target) {
        head = mid + 1;
        continue;
      }
      if (nums[mid] > target) {
        tail = mid;
        continue;
      }

      tail = mid;
    }

    return head;
  }
};
// @lc code=end
