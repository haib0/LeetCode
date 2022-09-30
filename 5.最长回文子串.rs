/*
 * @lc app=leetcode.cn id=5 lang=rust
 *
 * [5] 最长回文子串
 */

// @lc code=start
impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        fn is_palindrome(s: &str) -> bool {
            let mut s_l = s[..s.len() / 2].chars();
            let mut s_r = s[(s.len() + 1) / 2..].chars().rev();
            while let (Some(l), Some(r)) = (s_l.next(), s_r.next()) {
                if l != r {
                    return false;
                }
            }
            true
        }

        let mut result = String::from("");
        for length in (0..(s.len() + 1)).rev() {
            for index in 0..(s.len() - length + 1) {
                let piece = &s[index..index + length];

                if is_palindrome(piece) {
                    result = String::from(piece);
                    return result;
                }
            }
        }

        result
    }
}
// @lc code=end
