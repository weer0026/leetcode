package Binary_Gap

func binaryGap(N int) int {
	var (
		pos  = 0
		max  = 0
		meet = false // has been meet 1
	)
	for N > 0 {
		// rightest number is 1
		if N&1 == 1 {
			if meet && pos > max {
				max = pos
			}
			meet = true
			pos = 0
		}
		pos++
		N >>= 1
	}
	return max
}
