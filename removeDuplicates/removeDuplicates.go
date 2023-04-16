package removeduplicates

func removeDuplicates(nums []int) int {
	j := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			j++
			nums[j] = nums[i]
		}
	}
	j++
	nums[j] = nums[len(nums)-1]
	return j
}
