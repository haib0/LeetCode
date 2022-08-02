/*
 * @lc app=leetcode.cn id=9 lang=rust
 *
 * [9] 回文数
 */

// @lc code=start
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x >= 0 && x < 10 {
            return true;
        }
        if x < 0 || x % 10 == 0 {
            return false;
        }

        let mut x = x;
        let mut r = 0;
        while x > r {
            r = r * 10 + x % 10;
            x = x / 10;
        }
        return x == r || x == r / 10;
    }
}
// @lc code=end
