/*
 * @lc app=leetcode.cn id=13 lang=rust
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let s = s
            .replace("IV", "a")
            .replace("IX", "b")
            .replace("XL", "c")
            .replace("XC", "d")
            .replace("CD", "e")
            .replace("CM", "f");

        let mut result = 0;
        for c in s.chars() {
            result += match c {
                'I' => 1,
                'V' => 5,
                'X' => 10,
                'L' => 50,
                'C' => 100,
                'D' => 500,
                'M' => 1000,
                'a' => 4,
                'b' => 9,
                'c' => 40,
                'd' => 90,
                'e' => 400,
                'f' => 900,
                _ => 0,
            }
        }

        result
    }
}
// @lc code=end
