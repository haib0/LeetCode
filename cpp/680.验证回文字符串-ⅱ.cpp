/*
 * @lc app=leetcode.cn id=680 lang=cpp
 *
 * [680] 验证回文字符串 Ⅱ
 */
#include <string>
using namespace std;
// @lc code=start
class Solution {
public:
  bool isPal(string s, int left, int right) {
    while (left <= right) {
      if (s[left++] == s[right--]) {
        continue;
      }
      return false;
    }
    return true;
  }

  bool validPalindrome(string s) {
    int left = 0, right = s.size() - 1;
    while (left <= right) {
      if (s[left] == s[right]) {
        left++;
        right--;
        continue;
      }

      return isPal(s, left+1, right) || isPal(s, left, right-1);
    }

    return true;
  }
};
// @lc code=end
