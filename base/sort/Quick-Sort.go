package sort

func quickSort(data []int, left int, right int) {
	temp := data[left]
	p := left
	i, j := left, right

	// from right
	for j >= p && data[j] > temp {
		j--
	}
	data[p] = data[j]
	p = j
	// from left
	for i <= p && data[i] <= temp {
		i++
	}
	data[p] = data[i]
}

func QuickSort() {

}
