package leetcode

// 假设你有⼀个⻓度为 n 的数组，初始情况下所有的数字均为 0，你将会被给出 k 个更新的操作。
// 其中，每个操作会被表示为⼀个三元组：[startIndex, endIndex, inc]，你需要将⼦数组
// A[startIndex ... endIndex]（包括 startIndex 和 endIndex）增加 inc。
// 请你返回 k 次操作后的数组。
func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length)

	for _, update := range updates {
		start, end := update[0], update[1]
		inc := update[2]

		diff[start] += inc
		if end+1 < length {
			diff[end+1] -= inc
		}
	}

	nums := make([]int, length)
	tmp := 0
	for i, v := range diff {
		tmp += v
		nums[i] = tmp
	}
	return nums
}
