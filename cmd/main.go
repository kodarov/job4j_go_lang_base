package main

import (
	"fmt"

	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	nums := []int{1, 2, 1}
	res := base.Palindrome(nums)
	fmt.Println(res)

	var first *base.Node

	if first == nil {
		fmt.Println("nil pointer is used")
		first = &base.Node{
			Key:   "first",
			Value: "first",
		}
	}

	fmt.Printf("Node{key: %s, value: %s}\n", first.Key, first.Value)

}
