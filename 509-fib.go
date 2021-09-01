package leetcode

//斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//
//F(0) = 0，F(1) = 1
//F(n) = F(n - 1) + F(n - 2)，其中 n > 1
//给你 n ，请计算 F(n) 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/fibonacci-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func fib(n int) int {
	if n < 2 {
		return n
	}
	f := []int{0, 1}
	for i := 2; i <= n; i++ {
		f[0], f[1] = f[1], f[0]+f[1]
	}
	return f[1]
}
