package sliding_window

import "math"

func maxProfit(prices []int) int {
	min := math.MaxUint32
	rs := 0

	for _, price := range prices {
		if price > min {
			if price-min > rs {
				rs = price - min
			}
		} else {
			min = price
		}
	}

	return rs
}
