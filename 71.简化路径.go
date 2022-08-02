/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */
package leetcode

import (
	"strings"
)

// @lc code=start
func simplifyPath(path string) string {
	stack := []string{}

	for _, v := range strings.Split(path, "/") {
		switch v {
		case "", ".":
			// continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, v)
		}
	}

	return "/" + strings.Join(stack, "/")
}

// @lc code=end
