/*
 * @lc app=leetcode.cn id=58 lang=cpp
 *
 * [58] 最后一个单词的长度
 */
#include <string>

using namespace std;
// @lc code=start
class Solution
{
public:
    int lengthOfLastWord(string s)
    {
        int i = s.size() - 1;

        while (s[i] == ' ')
        {
            i--;
        }

        int result = 0;

        // 注意： i的判断需要在前，否则会错误！
        while (i >= 0 && s[i] != ' ')
        {
            result++;
            i--;
        }

        return result;
    }
};
// @lc code=end
