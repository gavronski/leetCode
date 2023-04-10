package searching

func BinarySearch(arr []int, value int) bool {
	low := 0
	high := len(arr) - 1

	var mid int

	for low <= high {
		mid = (low + high) / 2

		if value > arr[mid] {
			low = mid + 1
		} else if value < arr[mid] {
			high = mid - 1
		} else {
			return true
		}
	}

	return false
}

func BSRecursive(arr []int, value, high, low int) bool {
	if low > high {
		return false
	}

	mid := (low + high) / 2

	if value > arr[mid] {
		return BSRecursive(arr, value, high, mid+1)
	} else if value < arr[mid] {
		return BSRecursive(arr, value, mid-1, low)
	} else {
		return true
	}
}

func FindOccurence(arr []int, value int) int {
	low := 0
	high := len(arr) - 1
	occurence := -1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] == value {
			occurence = mid
			// first
			// high = mid - 1
			// last
			low = mid + 1
		} else if value > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return occurence
}
