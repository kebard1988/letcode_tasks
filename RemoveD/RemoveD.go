package removed

func removeDuplicates(nums []int) int {
	length, ne := len(nums), 1
	for i := 1; i < length; i++ {
		if nums[i-1] != nums[i] {
			nums[ne] = nums[i]
			ne++
		}
	}
	return ne
}