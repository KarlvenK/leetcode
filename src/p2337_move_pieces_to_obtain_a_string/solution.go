package p2337movepiecestoobtainastring

func canChange(start string, target string) bool {
	ps, pt := 0, 0
	n := len(start)
	for ps < n && pt < n {
		for ps < n && start[ps] == '_' {
			ps++
		}
		for pt < n && target[pt] == '_' {
			pt++
		}
		if ps >= n || pt >= n {
			break
		}
		if start[ps] != target[pt] {
			return false
		}
		if start[ps] == 'L' {
			if ps < pt {
				return false
			}
		} else {
			if ps > pt {
				return false
			}
		}
		ps++
		pt++
	}

	for ps < n {
		if start[ps] != '_' {
			return false
		}
		ps++
	}
	for pt < n  {
		if target[pt] != '_' {
			return false
		}
		pt++
	}


	return true
}
