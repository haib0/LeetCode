/*
 * @lc app=leetcode.cn id=1816 lang=cpp
 *
 * [1816] 截断句子
 */
#include <string>
using namespace std;
// @lc code=start
class Solution
{
public:
    string truncateSentence(string s, int k)
    {
        string rslt;
        int count = 1;
        for (size_t i = 0; i < s.length(); i++)
        {
            if (s[i] != ' ')
            {
                rslt += s[i];
            }
            else
            {
                if (count == k)
                {
                    return rslt;
                }
                rslt += ' ';
                count++;
            }
        }
        return rslt;
    }
};
// @lc code=end
