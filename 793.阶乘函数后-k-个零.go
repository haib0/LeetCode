/*
 * @lc app=leetcode.cn id=793 lang=golang
 *
 * [793] 阶乘函数后 K 个零
 */
package leetcode

// @lc code=start
func preimageSizeFZF(k int) int {
	left, right := left_bound(k), right_bound(k)
	return right - left

}

func left_bound(target int) int {
	left, right := 0, int(1e10)

	for left < right {
		mid := (left + right) / 2
		z := trailingZeroes0(mid)
		if z < target {
			left = mid + 1
			continue
		}
		if z >= target {
			right = mid
			continue
		}
	}

	return left

}

func right_bound(target int) int {
	left, right := 0, int(1e10)

	for left < right {
		mid := (left + right) / 2
		z := trailingZeroes0(mid)
		if z <= target {
			left = mid + 1
			continue
		}
		if z > target {
			right = mid
			continue
		}
	}

	return right

}

func trailingZeroes0(n int) int {
	ans := 0

	for n >= 5 {
		n = n / 5
		ans += n
	}

	return ans
}

// @lc code=end
