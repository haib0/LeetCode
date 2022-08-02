/*
 * @lc app=leetcode.cn id=125 lang=cpp
 *
 * [125] 验证回文串
 */
#include <cctype>
#include <string>
using namespace std;
// @lc code=start
class Solution {
public:
    bool isChar(char c) {
        if( (c < 'a' || c > 'a' + 25) && (c < 'A' || c > 'A' + 25) && (c < '0' || c > '9') ){
            return false;
        }


        return true;
    }
    bool isPalindrome(string s) {
        int left = 0, right = s.size()-1;

        while (right - left > 0) {
            if (!isChar(s[left])) {
            left++;
            continue;
            }
        
            if (!isChar(s[right])) {
            right--;
            continue;
            }


            if (tolower(s[left] )!= tolower(s[right])) {
                return false;
            }

            left++;
            right--;
        }

        return true;
    }
};
// @lc code=end

