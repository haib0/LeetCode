/*
 * @lc app=leetcode.cn id=911 lang=cpp
 *
 * [911] 在线选举
 */
#include <vector>
#include <unordered_map>

using namespace std;
// @lc code=start
class TopVotedCandidate
{
public:
    vector<int> tops;
    vector<int> times;
    TopVotedCandidate(vector<int> &persons, vector<int> &times)
    {
        this->times = times;

        unordered_map<int, int> topCount;
        int top = -1;
        for (int person : persons)
        {
            topCount[person]++;
            if (topCount[person] >= topCount[top])
            {
                top = person;
            }
            this->tops.emplace_back(top);
        }
    }

    int q(int t)
    {
        int head = 0;
        int tail = this->times.size() - 1;

        while (head < tail)
        {
            int mid = head + (tail - head + 1) / 2;
            if (this->times[mid] <= t)
            {
                head = mid;
            }
            else
            {
                tail = mid - 1;
            }
        }
        return this->tops[head];
    }
};

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * TopVotedCandidate* obj = new TopVotedCandidate(persons, times);
 * int param_1 = obj->q(t);
 */
// @lc code=end
