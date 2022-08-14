/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */
package leetcode

// @lc code=start
func isSubsequence(s string, t string) bool {
	m := make(map[rune][]int)

	for i, r := range t {
		if l, ok := m[r]; ok {
			m[r] = append(l, i)

		} else {
			m[r] = []int{i}
		}
	}

	j := 0
	for _, r := range s {
		if l, ok := m[r]; ok {
			// binary search j left bound in l
			left, right := 0, len(l)
			for left < right {
				mid := (left + right) / 2
				if l[mid] < j {
					left = mid + 1
				} else {
					right = mid
				}
			}

			if left == len(l) {
				return false
			}

			j = l[left] + 1
		} else {
			return false
		}
	}

	return true

}

// @lc code=end
