package base

func Mono(nums []int) bool {
	increasing, decreasing := true, true
	if len(nums) == 0 {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			increasing = false
		}
		if nums[i] < nums[i+1] {
			decreasing = false
		}
	}
	return increasing || decreasing
}
