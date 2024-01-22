package main

import "fmt"

func main() {

	var ages = [3]int{1, 2, 3}
	fmt.Println(ages, "and the length is", len(ages))

	names := [4]string{"a"}
	fmt.Println(names[0], "length", len(names))

	// slices

	test := []string{"abebe", "kebede"}
	test = append(test, "alemu")
	fmt.Println(test)

	//rangeOne := names[1:2]
	//rangeTwo := names[1:]
	rangeThree := test[:3]

	fmt.Println(rangeThree, "length", len(rangeThree))
}
