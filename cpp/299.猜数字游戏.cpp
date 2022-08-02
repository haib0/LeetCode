/*
 * @lc app=leetcode.cn id=299 lang=cpp
 *
 * [299] 猜数字游戏
 */
#include <string>
#include <unordered_map>

using namespace std;
// @lc code=start
class Solution
{
public:
    string getHint(string secret, string guess)
    {
        unordered_map<char, int> secret_map;
        vector<char> guess_vector;

        int A = 0;
        for (int i = 0; i < secret.length(); i++)
        {
            if (secret[i] == guess[i])
            {
                A++;
                continue;
            }
            secret_map[secret[i]]++;
            guess_vector.push_back(guess[i]);
        }

        int B = 0;
        for (int i = 0; i < guess_vector.size(); i++)
        {
            if (secret_map[guess_vector[i]] > 0)
            {
                B++;
                secret_map[guess_vector[i]]--;
            }
        }

        return to_string(A) + "A" + to_string(B) + "B";
    }
};
// @lc code=end
