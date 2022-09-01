package sol

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	res := make([]int, 2)
	for idx := range nums {
		remain := target - nums[idx]
		if val, exist := hash[remain]; exist {
			if val > idx {
				return []int{idx, val}
			} else {
				return []int{val, idx}
			}
		}
		hash[nums[idx]] = idx
	}
	return res
}
