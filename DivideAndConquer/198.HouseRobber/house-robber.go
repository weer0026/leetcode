package _98_HouseRobber

// dp[i] = max(dp[i-2] + nums[i], dp[i-1])
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[l-1]
}

// 内存优化，i>2时仅使用dp[i-1],dp[i-2]的值，双变量即可
func rob2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}
	// dp[i-2]
	first := nums[0]
	// dp[i-1]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := second
		// dp[i] = max(dp[i-2] + nums[i], dp[i-1])
		// 转化成 dp[i] = max(first + nums[i], second)
		second = max(first+nums[i], second)
		first = tmp
	}
	return second
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
