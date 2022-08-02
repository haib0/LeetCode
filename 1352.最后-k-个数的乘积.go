/*
 * @lc app=leetcode.cn id=1352 lang=golang
 *
 * [1352] 最后 K 个数的乘积
 */
package leetcode

// @lc code=start
type ProductOfNumbers struct {
	Products []int
}

func Constructor1352() ProductOfNumbers {
	return ProductOfNumbers{}
}

func (this *ProductOfNumbers) Add(num int) {
	for i, v := range this.Products {
		this.Products[i] = v * num
	}
	this.Products = append(this.Products, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.Products)
	return this.Products[n-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
// @lc code=end
