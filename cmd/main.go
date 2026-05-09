package main

import (
	"fmt"

	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	cache := base.NewLruCache(3)

	cache.Put("hello1", "world1")
	cache.Put("hello2", "world2")
	cache.Put("hello3", "world3")
	cache.Put("hello4", "world4")
	res := cache.Get("hello1")

	fmt.Println(res)

}
