/*
 * @lc app=leetcode.cn id=1 lang=rust
 *
 * [1] 两数之和
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = HashMap::with_capacity(nums.len());

        for (i, &value) in nums.iter().enumerate() {
            if let Some(&k) = map.get(&(target - value)) {
                if i < k {
                    return vec![i as i32, k as i32];
                } else if k < i {
                    return vec![k as i32, i as i32];
                }
            }
            map.insert(value, i);
        }
        panic!()
    }
}
// @lc code=end
