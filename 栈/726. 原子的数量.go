package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func countOfAtoms(formula string) string {
	i, n := 0, len(formula)

	parseAtom := func() string {
		start := i
		i++
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++
		}
		return formula[start:i]
	}

	parseNum := func() int {
		num := 0

		if i == n || !unicode.IsDigit(rune(formula[i])) {
			return 1
		}
		for ; i < n && unicode.IsDigit(rune(formula[i])); i++ {
			num = num*10 + int(formula[i]-'0')
		}
		return num
	}

	var stk []map[string]int
	stk = append(stk, make(map[string]int, 0))

	for i < n {
		if ch := formula[i]; ch == '(' {
			i++
			stk = append(stk, map[string]int{})
		} else if ch == ')' {
			i++
			num := parseNum()
			atomNum := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			for atom, v := range atomNum {
				stk[len(stk)-1][atom] += v * num
			}
		} else {
			atom := parseAtom()
			num := parseNum()
			stk[len(stk)-1][atom] += num
		}
	}

	atomNum := stk[0]
	type pair struct {
		atom string
		num  int
	}
	pairs := make([]pair, 0, len(atomNum))

	for k, v := range atomNum {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].atom < pairs[j].atom
	})
	var ans []byte
	for _, p := range pairs {
		ans = append(ans, []byte(p.atom)...)
		if p.num > 1 {
			ans = append(ans, strconv.Itoa(p.num)...)
		}
	}
	return string(ans)
}

func main() {
	ask := "K4(ON(SO3)2)2"
	fmt.Println(countOfAtoms(ask))
}
