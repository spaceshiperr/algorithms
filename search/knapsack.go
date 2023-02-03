package search

func GetKnapsackMaxProfit(m int, w, p []int) int {
	n := len(w)
	m++

	values := make([][]int, n)

	for i := 0; i < n; i++ {
		values[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				values[i][j] = 0
			} else if w[i] <= j {
				v1 := values[i-1][j]
				v2 := p[i] + values[i-1][j-w[i]]

				if v1 > v2 {
					values[i][j] = v1
				} else {
					values[i][j] = v2
				}
			} else {
				values[i][j] = values[i-1][j]
			}
		}
	}

	return values[n-1][m-1]
}
