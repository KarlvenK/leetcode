package search

const (
	eps = 1e-7
)

func judgePoint24(cards []int) bool {
	card := make([]float64, 0)
	for i := 0; i < 4; i++ {
		card = append(card, float64(cards[i]))
	}
	return dfs(card)
}

func dfs(cards []float64) bool {
	sz := len(cards)
	if sz == 0 {
		return false
	}
	if sz == 1 {
		if abs(cards[0]-24.0) < eps {
			return true
		}
		return false
	}
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			if i != j {
				list := make([]float64, 0)
				for t := 0; t < sz; t++ {
					if t != i && t != j {
						list = append(list, cards[t])
					}
				}
				num := 0.0
				for k := 0; k < 4; k++ {
					switch k {
					case 0:
						num = cards[i] + cards[j]
					case 1:
						num = cards[i] - cards[j]
					case 2:
						num = cards[i] * cards[j]
					case 3:
						if cards[j] < eps {
							continue
						} else {
							num = cards[i] / cards[j]
						}
					}
					list = append(list, num)
					if dfs(list) == true {
						return true
					}
					list = list[0 : len(list)-1]
				}
			}
		}
	}
	return false
}
func abs(x float64) float64 {
	if x < 0 {
		return x * -1.0
	}
	return x
}
