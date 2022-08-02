/*
 * @lc app=leetcode.cn id=34 lang=cpp
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
#include <vector>
using namespace std;
// @lc code=start
class Solution {
public:
  vector<int> searchRange(vector<int> &nums, int target) {
    int head = 0, tail = nums.size();

    while (head < tail) {
      int i = (tail + head) / 2;

      if (nums[i] < target) {
        head = i + 1;
        continue;
      }
      if (nums[i] > target) {
        tail = i;
        continue;
      }

      int left = i, right = i;
      while (left >= 0 && nums[left] == target) {
        left--;
      }
      while (right < nums.size() && nums[right] == target) {
        right++;
      }
      return vector<int>{++left, --right};
    }

    return vector<int>{-1, -1};
  }
};
// @lc code=end
