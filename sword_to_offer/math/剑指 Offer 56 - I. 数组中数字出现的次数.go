package math

func singleNumbers(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor = xor ^ v
	}
	lowbit := xor & (-xor)
	var x, y int
	for _, v := range nums {
		if v&lowbit == 0 {
			x = x ^ v
		} else {
			y = y ^ v
		}
	}
	return []int{x, y}
}

/*
	题面：一个[]int,其中除了某两个数字，其它数字都出现了两次，求出这两个数字。

	solution：
			设两个不一样的数字分别为x, y
			由于相同的数字xor结果为0， 所以我们可以快速得出x xor y
			但是需要我们分别得出x和y是多少
			所以我们可以考虑将x 和 y 分成两组，相同的数字在同一组，那么每一组都进行xor，就可以得出x和y
			所以我们要做的就是将x和y分到不同的两组
			由于x和y不相同，那么二进制形式的x和y至少有一位不相同，那么我们可以用lowbit公式得出最右的1的位置，并把只有这一位为1的数记为m（lowbit就是直接诶得出这个m的）
			x和y的二进制形式在这位不相同
			那么我们可以根据这一位为0还是1将数字分成两组，相同的数字必定在一组
			这样就达成而来将x和y分到不同的两组，最后再分别求xor就可以得出x和y是多少了
*/
