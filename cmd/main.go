package main

import (
	"fmt"

	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	nums := []int{1, 2, 1}
	res := base.Palindrome(nums)
	fmt.Println(res)
}
