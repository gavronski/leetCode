package sorting

func BubbleSort(items []int) []int {
	len := len(items)

	for i := 0; i < len; i++ {
		swapped := false
		for j := 0; j < len-i-1; j++ {
			if items[j] > items[j+1] {
				tmp := items[j]
				items[j] = items[j+1]
				items[j+1] = tmp

				swapped = true
			}
		}

		if !swapped {
			return items
		}
	}

	return items
}

func SelectionSort(items []int) []int {
	len := len(items)

	for i := 0; i < len; i++ {
		min := i

		for j := i + 1; j < len; j++ {
			if items[j] < items[min] {
				min = j
			}
		}

		if min != i {
			tmp := items[min]
			items[min] = items[i]
			items[i] = tmp
		}
	}

	return items
}

func InsertionSort(items []int) []int {
	len := len(items)

	for i := 1; i < len; i++ {
		start := items[i]
		j := i - 1

		for j >= 0 && items[j] > start {
			items[j+1] = items[j]
			j--
		}

		items[j+1] = start
	}

	return items
}

func MergeSort(items []int) []int {
	len := len(items)

	if len == 1 {
		return items
	}

	first := MergeSort(items[:len/2])
	second := MergeSort(items[len/2:])

	return merge(first, second)
}

func merge(a, b []int) []int {
	i := 0
	j := 0

	sorted := []int{}

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		sorted = append(sorted, a[i])
	}

	for ; j < len(b); j++ {
		sorted = append(sorted, b[j])
	}

	return sorted
}
