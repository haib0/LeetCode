/*
 * @lc app=leetcode.cn id=633 lang=cpp
 *
 * [633] 平方数之和
 */

#include "corecrt_math.h"
#include <vcruntime.h>

// @lc code=start
class Solution {
public:
  bool judgeSquareSum(int c) {
    size_t i = 0;
    size_t j = sqrt(c);
    while (i <= j) {
      size_t s = i * i + j * j;

      if (s == c) {
        return true;
      }

      if (s < c) {
        i++;
      } else {
        j--;
      }
    }
    return false;
  }
};
// @lc code=end
