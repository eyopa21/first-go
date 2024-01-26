package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	test := "Hello world"
	fmt.Println(strings.Contains(test, "Hello"))
	fmt.Println(strings.ReplaceAll(test, "Hello", "Greetings"))
	fmt.Println(len("hello world"))

	ages := []int{1, 2, 5, 3, 9, 6}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 0)

	fmt.Println(index)

}
