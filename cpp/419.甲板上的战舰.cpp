/*
 * @lc app=leetcode.cn id=419 lang=cpp
 *
 * [419] 甲板上的战舰
 */
#include <vector>
using namespace std;

// @lc code=start
class Solution
{
public:
    int countBattleships(vector<vector<char>> &board)
    {
        int ans = 0;

        for (int i = 0; i < board.size(); i++)
        {
            for (int j = 0; j < board[0].size(); j++)
            {
                if (board[i][j] == 'X')
                {
                    if (i > 0 && board[i - 1][j] == 'X')
                    {
                        continue;
                    }
                    if (j > 0 && board[i][j - 1] == 'X')
                    {
                        continue;
                    }
                    ans++;
                }
            }
        }
        return ans;
    }
};
// @lc code=end
