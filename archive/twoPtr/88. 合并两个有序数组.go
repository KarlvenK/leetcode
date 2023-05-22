package twoPtr

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := n + m - 1; i >= n; i-- {
		nums1[i] = nums1[i-n]
	}
	p1, p2 := n, 0
	inx := 0
	//fmt.Println(nums1, nums2)
	for inx < m+n {
		if p1 == m+n || p2 == n {
			break
		}
		if nums1[p1] > nums2[p2] {
			nums1[inx] = nums2[p2]
			p2++
			inx++
		} else {
			nums1[inx] = nums1[p1]
			p1++
			inx++
		}
	}

	for p1 < m+n {
		nums1[inx] = nums1[p1]
		p1++
		inx++
	}
	for p2 < n {
		nums1[inx] = nums2[p2]
		inx++
		p2++
	}
}
