package leetcode

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindrome-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	res, rx := 0, x
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res == rx
}
