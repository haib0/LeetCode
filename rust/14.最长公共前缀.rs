/*
 * @lc app=leetcode.cn id=14 lang=rust
 *
 * [14] 最长公共前缀
 */

// @lc code=start
impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let mut common_prefix = String::new();
        let mut strs = strs;
        let s0 = strs.pop().unwrap();
        for i in 1..s0.len() {
            let prefix = String::from(s0.split_at(i).0);
            for s in &strs {
                if s.starts_with(&prefix) {
                    continue;
                } else {
                    return common_prefix;
                }
            }
            common_prefix = prefix.to_string();
        }
    
        for s in &strs {
            if s.starts_with(&s0) {
                continue;
            } else {
                return common_prefix;
            }
        }
        return s0.to_string();
    
    }
}
// @lc code=end

