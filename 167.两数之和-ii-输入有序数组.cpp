/*
 * @lc app=leetcode.cn id=167 lang=cpp
 *
 * [167] 两数之和 II - 输入有序数组
 */
#include <vector>

using namespace std;

// @lc code=start
class Solution {
public:
  vector<int> twoSum(vector<int> &numbers, int target) {
    int i = 0, j = numbers.size() - 1;
    while (i < j) {
      int sum = numbers[i] + numbers[j];
      if (sum == target) {
        return {i + 1, j + 1};
      }
      if (sum < target) {
        i++;
      } else {
        j--;
      }
    }
    return {};
  }
};
// @lc code=end
