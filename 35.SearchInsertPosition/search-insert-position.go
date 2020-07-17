package SearchInsertPosition

// 二分 O(logN)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)>>1 + left // 位运算增加计算效率，拆分计算防止left,right过大导致整形溢出，本质还是(left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 暴力遍历 O(n)
func searchInsert2(nums []int, target int) int {
	if nums == nil {
		return 0
	}
	for i, v := range nums {
		if v == target {
			return i
		}

		if v > target {
			return i
		}

		if i == len(nums)-1 {
			return i + 1
		}
	}
	return 0
}
