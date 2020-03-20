package MaximumSubarray

func maxSubArray(nums []int) int {
	dp_sum, max := nums[0], nums[0]

	for i := 0; i < len(nums); i++ {
		if dp_sum > 0 {
			dp_sum += nums[i]
		} else {
			dp_sum = nums[i]
		}

		if dp_sum > max {
			max = dp_sum
		}
	}
	return max
}
