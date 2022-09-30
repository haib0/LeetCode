/*
 * @lc app=leetcode.cn id=1034 lang=cpp
 *
 * [1034] 边框着色
 */
#include <vector>
using namespace std;
// @lc code=start
class Solution
{
public:
    vector<vector<int>> colorBorder(vector<vector<int>> &grid, int row, int col, int color)
    {
        int oColor = grid[row][col];
        vector<vector<bool>> visited(grid.size(), vector<bool>(grid[0].size(), false));
        vector<pair<int, int>> borders;
        dfs(grid, {row, col}, visited, borders, oColor);
        for (auto &b : borders)
        {
            int x = b.first;
            int y = b.second;
            grid[x][y] = color;
        }
        return grid;
    }

    void dfs(vector<vector<int>> &grid, pair<int, int> point, vector<vector<bool>> &visited, vector<pair<int, int>> &borders, int oColor)
    {
        visited[point.first][point.second] = true;

        int m = grid.size(), n = grid[0].size();
        bool isBorder = false;
        pair<int, int> direcs[4] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

        for (auto &direc : direcs)
        {
            int x = point.first + direc.first, y = point.second + direc.second;
            if (x < 0 || y < 0 || x == m || y == n || grid[x][y] != oColor)
            {
                isBorder = true;
            }
            else if (!visited[x][y])
            {
                dfs(grid, {x, y}, visited, borders, oColor);
            }
        }
        if (isBorder)
        {
            borders.emplace_back(point);
        }
    }
};
// @lc code=end
