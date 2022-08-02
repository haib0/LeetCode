/*
 * @lc app=leetcode.cn id=438 lang=cpp
 *
 * [438] 找到字符串中所有字母异位词
 */
#include <string>
#include <vector>
using namespace std;
// @lc code=start
class Solution {
public:
  vector<int> findAnagrams(string s, string p) {
    int sLen = s.size(), pLen = p.size();
    
    if (sLen < pLen) {
      return vector<int>();
    }

    vector<int> results;

    vector<int> sCount(26);
    vector<int> pCount(26);

    for (int i = 0; i < pLen; i++) {
      ++sCount[s[i] - 'a'];
      ++pCount[p[i] - 'a'];
    }

    if (sCount == pCount) {
      results.emplace_back(0);
    }

    for (int i = 0; i < sLen - pLen; i++) {
      --sCount[s[i] - 'a'];
      ++sCount[s[i + pLen] - 'a'];
      
      if (sCount == pCount) {
        results.emplace_back(i + 1);
      }
    }

    return results;
  }
};
// @lc code=end
