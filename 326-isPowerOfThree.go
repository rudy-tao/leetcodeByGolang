package leetcode

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n%3 == 0 {
		n = n / 3
	}
	return n == 1
}
