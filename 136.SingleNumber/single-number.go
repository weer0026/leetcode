package _36_SingleNumber

// 异或： a^a相同为零，a^0=a
func singleNumber(nums []int) int {
	num := 0
	for _, v := range nums {
		num = num ^ v
	}
	return num
}
