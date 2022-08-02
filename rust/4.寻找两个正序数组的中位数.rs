/*
 * @lc app=leetcode.cn id=4 lang=rust
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let length = nums1.len() + nums2.len();
        let m_loc = length / 2;
        let (mut i, mut j, mut m1, mut m2) = (0, 0, 0, 0);
        loop {
            if i == nums1.len() {
                if i + j + 1 == m_loc {
                    m1 = nums2[j];
                }
                if i + j + 1 == m_loc + 1 {
                    m2 = nums2[j];
                    break;
                }
                j += 1;
                continue;
            }
            if j == nums2.len() {
                if i + j + 1 == m_loc {
                    m1 = nums1[i];
                }
                if i + j + 1 == m_loc + 1 {
                    m2 = nums1[i];
                    break;
                }
                i += 1;
                continue;
            }
    
            let n1 = nums1[i];
            let n2 = nums2[j];
            if n1 < n2 {
                if i + j + 1 == m_loc {
                    m1 = n1;
                }
                if i + j + 1 == m_loc + 1 {
                    m2 = n1;
                    break;
                }
                i += 1;
            } else {
                if i + j + 1 == m_loc {
                    m1 = n2;
                }
                if i + j + 1 == m_loc + 1 {
                    m2 = n2;
                    break;
                }
                j += 1;
            }
        }
        if length % 2 == 0 {
            return (m1 + m2) as f64 / 2.0;
        } else {
            return m2 as f64;
        }
    }
    
}
// @lc code=end
