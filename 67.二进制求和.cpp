/*
 * @lc app=leetcode.cn id=67 lang=cpp
 *
 * [67] 二进制求和
 */
#include <string>
using namespace std;
// @lc code=start
class Solution
{
public:
    string addBinary(string a, string b)
    {
        string ans = "";

        int la = a.size();
        int lb = b.size();
        if (la < lb)
        {
            string prev = "";
            for (int t = 0; t < lb - la; t++)
            {
                prev += '0';
            }
            a = prev + a;
        }
        if (lb < la)
        {
            string prev = "";
            for (int t = 0; t < la - lb; t++)
            {
                prev += '0';
            }
            b = prev + b;
        }

        char flag = '0';
        int i;
        for (i = a.size() - 1; i >= 0; i--)
        {
            char ca = a[i], cb = b[i];
            // '0'=48, '1'=49
            switch (ca + cb + flag)
            {
            case '0' + '0' + '0':
                ans = '0' + ans;
                break;
            case '0' + '0' + '1':
                ans = '1' + ans;
                flag = '0';
                break;
            case '0' + '1' + '1':
                ans = '0' + ans;
                flag = '1';
                break;
            case '1' + '1' + '1':
                ans = '1' + ans;
                flag = '1';
                break;
            default:
                break;
            }
        }
        if (flag == '1')
        {
            ans = flag + ans;
        }
        return ans;
    }
};
// @lc code=end
