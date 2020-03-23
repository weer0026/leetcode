package MajorityElement

/**
hash
*/

func majorityElement(nums []int) int {
	expected := len(nums) / 2
	numsCount := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := numsCount[nums[i]]; ok {
			numsCount[nums[i]]++
		} else {
			numsCount[nums[i]] = 1
		}
	}

	for k, v := range numsCount {
		if (v >= expected) {
			return k
		}
	}
	return 0
}

/**
	boyer-moore voting algorithm
 */
func majorityElementBM(nums []int) int {
	count := 1
	cand := nums[0]

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			cand = nums[i]
		}
		if cand == nums[i] {
			count++
			continue
		}
		count--
	}
	return cand;
}
