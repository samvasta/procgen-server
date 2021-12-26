package common

func Min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}

func MinF(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}
