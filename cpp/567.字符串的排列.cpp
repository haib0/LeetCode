/*
 * @lc app=leetcode.cn id=567 lang=cpp
 *
 * [567] 字符串的排列
 */
#include <iostream>
#include <string>
#include <unordered_map>
using namespace std;
// @lc code=start
class Solution {
public:
  bool checkInclusion(string s1, string s2) {
    unordered_map<char, int> mp;
    for (auto &c : s1) {
      mp[c]++;
    }
    if (s2.size() < s1.size()) {
      return false;
    }
    for (int i = 0; i < s2.size() - s1.size() + 1; i++) {
      auto mpi(mp);
      bool ok = false;

      for (int j = i; j < i + s1.size(); j++) {
        // cout << s2[j] << " ";
        if (mpi.find(s2[j]) != mpi.end()) {
          mpi[s2[j]]--;
          //   cout << s2[j] << " ";
        } else {
          break;
        }

        ok = true;
      }

      if (!ok) {
        continue;
      }

      for (auto &x : mpi) {
        if (x.second != 0) {
          ok = false;
          break;
        }
      }

      if (ok) {
        return true;
        ;
      }
    }

    return false;
  }
};
// @lc code=end
