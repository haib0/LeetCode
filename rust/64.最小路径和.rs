/*
 * @lc app=leetcode.cn id=64 lang=rust
 *
 * [64] 最小路径和
 */

// @lc code=start
impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut dp: Vec<Vec<i32>> = vec![vec![grid[0][0];n];m];
        for i in 1..m {
            dp[i][0] = grid[i][0] + dp[i-1][0];
        }
        for j in 1..n {
            dp[0][j] = grid[0][j] + dp[0][j-1];
        }
        for i in 1..m {
            for j in 1..n {
                let a = dp[i-1][j];
                let b = dp[i][j-1];
                dp[i][j] = grid[i][j] + {if a < b {a} else {b}};
            }
        }
        return dp[m-1][n-1];
    }
}
// @lc code=end

