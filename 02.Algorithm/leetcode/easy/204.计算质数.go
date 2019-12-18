package easy

import "math"

/*
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	c := 0
	lastPrime := 2
	for num := 2; num < n; num++ {
		if isPrimesCacheVersion(num, lastPrime) {
			lastPrime = num
			c += 1
		}
	}

	return c
}

func isPrimesCacheVersion(n, lastPrime int) bool {
	sqrt := int(math.Ceil(math.Sqrt(float64(n))))
	var i int
	if lastPrime > sqrt {
		i = 2
	} else {
		i = lastPrime
	}
	for ; i <= sqrt; i++ {
		if n%i == 0 && i != n {
			return false
		}
	}
	return true
}
