package _0001_two_sum

func twoSum(nums []int, target int) []int {

	for i := range nums {
		for j := range nums {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
