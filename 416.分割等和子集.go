/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package leetcode

// @lc code=start
func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	// the ans is false when sum is odd
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2

	var n = len(nums)
	// dp[i][j]=true means nums[:i] can gene a new array that sums eq j
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	for i := 1; i < n; i++ {
		for s := 1; s <= sum; s++ {
			if s < nums[i] {
				dp[i][s] = dp[i-1][s]
			} else {
				dp[i][s] = dp[i-1][s-nums[i]] || dp[i-1][s]
			}
		}
	}

	return dp[n-1][sum]
}

// @lc code=end
