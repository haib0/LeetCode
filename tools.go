package leetcode

var AbsInt = func(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

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

var PowInt = func(x, a int) int {
	y := 1
	for i := 0; i < a; i++ {
		y = y * x
	}

	return y
}
