/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
package leetcode

// @lc code=start
func myAtoi(s string) int {
	MAX := 2147483647
	MIN := -2147483648

	flag := 1
	flagHappen := false
	trim0 := false
	l := []rune{}

	for _, c := range s {
		if c == ' ' {
			if len(l) > 0 || flagHappen {
				break
			}
			if trim0 {
				return 0
			}
			continue
		}
		if c == '0' && len(l) == 0 {
			trim0 = true
			continue
		}
		if c == '-' {
			if flagHappen || trim0 || len(l) > 0 {
				break
			}
			flag = -1
			flagHappen = true
			continue

		}
		if c == '+' {
			if flagHappen || trim0 || len(l) > 0 {
				break
			}
			flag = 1
			flagHappen = true
			continue
		}
		if c < '0' || c > '9' {
			break
		}
		l = append(l, c)
	}
	if len(l) > 10 {
		if flag == 1 {
			return MAX
		} else {
			return MIN
		}
	}

	r := 0
	x := 1
	for i := len(l) - 1; i >= 0; i-- {
		r += (int(l[i]) - 48) * x
		if r > MAX && flag == 1 {
			return MAX
		}
		if r*flag < MIN && flag == -1 {
			return MIN
		}
		x *= 10
	}
	return r * flag
}

// @lc code=end
