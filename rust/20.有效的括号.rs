/*
 * @lc app=leetcode.cn id=20 lang=rust
 *
 * [20] 有效的括号
 */

// @lc code=start
impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack: Vec<char> = vec![];

        let match_closures = |x| -> char {
            match x {
                ']' => '[',
                '}' => '{',
                ')' => '(',
                _ => '1',
            }
        };

        for c in s.chars() {
            if stack.is_empty() {
                stack.push(c);
            } else {
                let a = *stack.last().unwrap();
                if a.eq(&match_closures(c)) {
                    stack.pop();
                } else {
                    stack.push(c);
                }
            }
        }

        if stack.is_empty() {
            true
        } else {
            false
        }
    }
}
// @lc code=end
