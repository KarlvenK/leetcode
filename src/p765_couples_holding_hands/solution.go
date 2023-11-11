package p765_couples_holding_hands

type unionSet struct {
	n   int
	par []int
}

func newUnionSet(n int) unionSet {
	par := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
	}
	return unionSet{
		n:   n,
		par: par,
	}
}

func (u unionSet) find(x int) int {
	if x == u.par[x] {
		return x
	}
	u.par[x] = u.find(u.par[x])
	return u.par[x]
}

func (u unionSet) union(x, y int) {
	pa := u.find(x)
	pb := u.find(y)
	u.par[pa] = pb
}

func (u unionSet) count() int {
	ret := 0
	for i := 0; i < u.n; i++ {
		if i == u.find(i) {
			ret++
		}
	}
	return ret
}

func minSwapsCouples(row []int) int {
	n := len(row) / 2
	set := newUnionSet(n)

	for i := 0; i < n*2; i += 2 {
		set.union(row[i]/2, row[i+1]/2)
	}
	return n - set.count()
}
