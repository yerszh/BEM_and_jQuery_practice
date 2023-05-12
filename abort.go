package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}

	medianIndex := len(nums) / 2

	return nums[medianIndex]
}
