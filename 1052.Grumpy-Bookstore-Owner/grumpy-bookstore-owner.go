package _052_Grumpy_Bookstore_Owner

func maxSatisfied(customers []int, grumpy []int, x int) int {
	// 极限情况

	normal, more, length := 0, 0, len(customers)
	// 计算正常客人数
	for i := 0; i < length; i++ {
		if grumpy[i] == 0 {
			normal += customers[i]
		}
	}

	// 计算技巧挽留客人数 滑动窗口
	// 第一个窗口挽留人数 [0, x - 1]
	for i := 0; i < x; i++ {
		more += customers[i] * grumpy[i]
	}

	maxMore := more
	for i := x; i < length; i++ {
		// 滑动窗口
		// 减去离开的 i-x, 增加滑动进去的 i
		more = more - customers[i-x]*grumpy[i-x] + customers[i]*grumpy[i]
		maxMore = max(maxMore, more)
	}
	return normal + maxMore
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
