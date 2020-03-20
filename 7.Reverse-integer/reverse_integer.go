package reverse_integer



func reverse(x int) int {
	MAX := 0x80000000
	MIN := 0x7FFFFFFF
	sum := 0

	for {
		left := x / 10
		last := x % 10
		x = left
		sum = sum * 10 + last

		if 0 == left {
			break
		}
	}

	if sum < -MIN || sum > MAX {
		return 0
	}
	return sum
}
