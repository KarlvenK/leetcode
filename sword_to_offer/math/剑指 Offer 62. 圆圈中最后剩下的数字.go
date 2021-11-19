package math

func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n
}

/*
题目中的要求可以表述为：给定一个长度为 n 的序列，每次向后数 m 个元素并删除，那么最终留下的是第几个元素？

设答案为f(n,m)
那么第一次会让 m % n号的元素出列，那么就是要在剩下n - 1继续进行操作
设最后的生存者为f(n - 1, m)
由于每次都是从出列者的下一位开始记为0号位
那么f(n - 1,m)得出的就是m % n之后的offset
所以f(n,m)可以推出为 （f(n - 1, m) +　ｍ　％　ｎ）　％　ｎ
*/
