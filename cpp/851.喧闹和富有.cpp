/*
 * @lc app=leetcode.cn id=851 lang=cpp
 *
 * [851] 喧闹和富有
 */
#include <vector>
#include <unordered_map>
using namespace std;

// @lc code=start
class Solution
{
public:
    vector<int> loudAndRich(vector<vector<int>> &richer, vector<int> &quiet)
    {
        int n = quiet.size();

        unordered_map<int, vector<int>> rMap;
        for (auto &p : richer)
        {
            rMap[p[1]].push_back(p[0]);
        }

        vector<int> result = vector<int>(n);
        for (int i = 0; i < n; i++)
        {
            int minQ = quiet[i];
            int minP = i;

            vector<int> queue = rMap[i];
            // cout <<"QUEUE" << queue.size() <<endl;
            while (queue.size() > 0)
            {
                int p = queue[0];
                queue.erase(queue.begin());
                // cout << p <<endl;
                if (p == i)
                {
                    continue;
                }
                if (p < i && quiet[result[p]] < minQ)
                {
                    minQ =  quiet[result[p]];
                    minP =  result[p];
                    continue;
                }
                if (quiet[p] < minQ)
                {
                    minQ = quiet[p];
                    minP = p;
                }
                for (auto &nP : rMap[p])
                {
                    queue.push_back(nP);
                }
            }
            result[i] = minP;
            // cout<< "MINP" << minP <<endl;
        }
        return result;
    }
};
// @lc code=end
