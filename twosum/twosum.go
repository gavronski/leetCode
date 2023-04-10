package twosum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

func TwoSum(nums []int, target int) []int {
	var output []int
	keyMap := make(map[int]int)
	for k, v := range nums {
		keyMap[v] = k
	}

	for k, v := range nums {
		r := target - v
		// check if map contains value for given key
		_, ok := keyMap[r]

		if ok && keyMap[r] != k {
			output = append(output, k, keyMap[r])
			break
		}
	}
	return output
}
