/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 */
package leetcode

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i1, i2 := 0, 0
	var ans [][]int
	for i1 < len(firstList) && i2 < len(secondList) {
		a1, a2 := firstList[i1][0], firstList[i1][1]
		b1, b2 := secondList[i2][0], secondList[i2][1]

		if a1 <= b2 && b1 <= a2 {
			l, r := a1, b2
			if a1 < b1 {
				l = b1
			}
			if a2 < b2 {
				r = a2
			}
			ans = append(ans, []int{l, r})
		}

		if a2 < b2 {
			i1++
		} else {
			i2++
		}
	}
	return ans
}

// @lc code=end
