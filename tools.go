package leetcode

var minInt = func(nums ...int) int {
	if len(nums) == 0 {
		panic(nil)
	}

	x := nums[0]
	for _, v := range nums[1:] {
		if v < x {
			x = v
		}
	}
	return x
}

var maxInt = func(nums ...int) int {
	if len(nums) == 0 {
		panic(nil)
	}

	x := nums[0]
	for _, v := range nums[1:] {
		if v > x {
			x = v
		}
	}

	return x
}
