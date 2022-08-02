/*
 * @lc app=leetcode.cn id=748 lang=cpp
 *
 * [748] 最短补全词
 */
#include <vector>
#include <string>
#include <unordered_map>

using namespace std;

// @lc code=start
class Solution
{
public:
    string shortestCompletingWord(string licensePlate, vector<string> &words)
    {
        int min_length = INT_MAX;
        string result;

        unordered_map<char, int> mp;
        for (auto &c : licensePlate)
        {
            if (isalpha(c))
            {
                c = tolower(c);
                mp[c]++;
            }
        }

        for (auto &word : words)
        {
            bool isOk = true;
            unordered_map<char, int> m = {};
            for (auto &c : word)
            {
                m[c]++;
            }

            for (auto &c : licensePlate)
            {
                if (m[c] < mp[c])
                {
                    isOk = false;
                    break;
                }
            }

            if (!isOk)
            {
                continue;
            }
            int length = word.size();
            if (length < min_length)
            {
                min_length = length;
                result = word;
            }
        }
        return result;
    }
};
// @lc code=end
