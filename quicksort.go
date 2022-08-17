package leetcode

func QuickSort[T any](s []T, fn func(i, j int) bool) {
	partition := func(lo, hi int) int {
		i, j := lo+1, hi
		for i <= j {
			for i < hi && fn(i, lo) {
				i++
			}
			for j > lo && !fn(j, lo) {
				j--
			}
			if i >= j {
				break
			}
			s[i], s[j] = s[j], s[i]
		}
		s[j], s[lo] = s[lo], s[j]
		return j
	}

	var sort func(lo, hi int)
	sort = func(lo, hi int) {
		if hi <= lo {
			return
		}
		p := partition(lo, hi)
		sort(lo, p-1)
		sort(p+1, hi)
	}

	sort(0, len(s)-1)
}
