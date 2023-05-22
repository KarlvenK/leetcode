package main

type NumArray struct {
	vals    []int
	sz      int
	bitTree []int
}

func lowbit(x int) int {
	return x & (-x)
}
func (numArr *NumArray) Add(x, k int) {
	x++
	for x <= numArr.sz {
		numArr.bitTree[x] += k
		x += lowbit(x)
	}
}
func (numArr *NumArray) GetSum(x int) int {
	x++
	ans := 0
	for x >= 1 {
		ans += numArr.bitTree[x]
		x -= lowbit(x)
	}
	return ans
}

func Constructor(nums []int) NumArray {
	ret := NumArray{
		vals: nums,
		sz:   len(nums),
	}

	ret.bitTree = make([]int, ret.sz+1)

	for i := 0; i < ret.sz; i++ {
		ret.Add(i, nums[i])
	}
	return ret
}

func (numArr *NumArray) Update(index int, val int) {
	delta := val - numArr.vals[index]
	numArr.vals[index] = val
	numArr.Add(index, delta)
}

func (numArr *NumArray) SumRange(left int, right int) int {
	return numArr.GetSum(right) - numArr.GetSum(left-1)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
