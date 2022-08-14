/*
 * @lc app=leetcode.cn id=391 lang=golang
 *
 * [391] 完美矩形
 */
package leetcode

import "fmt"

// @lc code=start
func isRectangleCover(rectangles [][]int) bool {
	type point struct {
		x, y int
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

	minX, minY, maxX, maxY := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	mp := make(map[point]int)
	area := 0
	for _, v := range rectangles {
		minX = minInt(minX, v[0])
		minY = minInt(minY, v[1])
		maxX = maxInt(maxX, v[2])
		maxY = maxInt(maxY, v[3])

		area += (v[2] - v[0]) * (v[3] - v[1])

		mp[point{x: v[0], y: v[1]}]++
		mp[point{x: v[0], y: v[3]}]++
		mp[point{x: v[2], y: v[1]}]++
		mp[point{x: v[2], y: v[3]}]++
	}

	if area != (maxX-minX)*(maxY-minY) {
		fmt.Println(area)
		return false
	}

	p4 := []point{{minX, minY}, {minX, maxY}, {maxX, minY}, {maxX, maxY}}
	for _, p := range p4 {
		if _, ok := mp[p]; ok {
			delete(mp, p)
		} else {
			return false
		}
	}

	for _, n := range mp {
		if (n != 2) && (n != 4) {
			return false
		}
	}

	return true
}

// @lc code=end
