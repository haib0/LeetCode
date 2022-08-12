/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */
package leetcode

// @lc code=start

func GeneNature() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func countPrimesCo(n int) int {
	var ans int
	ch := GeneNature()
	for {
		prime := <-ch
		if prime >= n {
			return ans
		}
		ans++
		ch = PrimeFilter(ch, prime)
	}
}

func countPrimes(n int) int {
	notPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !notPrime[i] {
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}

	ans := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			ans++
		}
	}
	return ans
}

// @lc code=end
