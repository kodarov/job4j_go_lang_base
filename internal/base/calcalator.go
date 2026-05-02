package base

func Add(first, second int) int {
	return first + second
}

func Max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func Palindrome(nums []int) bool {
	size := len(nums)
	mid := size / 2

	for i := 0; i < mid; i++ {
		if nums[i] != nums[size-i-1] {
			return false
		}
	}
	return true
}
