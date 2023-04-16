package searchinsertposition

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.

func SearchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	var mid int

	for low <= high {
		mid = (low + high) / 2

		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return low
}
