/*
 * @lc app=leetcode.cn id=3 lang=rust
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        use std::collections::HashMap;
        use std::convert::TryInto;
    
        let mut head = 0;
        let mut tail = 0;
    
        let s: Vec<char> = s.chars().collect();
        let mut max_length = 0;
    
        while tail < s.len() {
            let mut is_dup = false;
            let mut map: HashMap<char, bool> = HashMap::new();
            let mut i = head;
            while i <= tail {
                if let Some(_) = map.get(&s[i]) {
                    is_dup = true;
                }
                map.insert(s[i], true);
                i += 1;
            }
            if is_dup {
                head += 1;
            } else {
                max_length = if max_length < tail - head + 1 {
                    tail - head + 1
                } else {
                    max_length
                };
                tail += 1;
            }
        }
    
        max_length.try_into().unwrap()
    }
}
// @lc code=end
