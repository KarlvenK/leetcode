package main

func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		ans++
		num &= num - 1
	}
	return ans
}

func main() {

}

/*
题目：给你一个数num，求它二进制形式下，1的个数

要点：
num & (num - 1) 等于将最右的1置零
e.g.
num    = 	  1xxxxx10000
num - 1= 	  1xxxxx01111
num & (num - 1) = 1xxxxx00000
*/
