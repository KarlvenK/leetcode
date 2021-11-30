package interview
func convertInteger(A int, B int) int {
	c := int32(A ^ B)
	ans := 0
	for c != 0 {
		ans++
		c = c & (c - 1)
	}
	return ans
}
