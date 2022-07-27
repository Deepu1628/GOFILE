package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	//**********************
	var array6 [20]int
	array6[0] = 30
	array6[1] = 4
	fmt.Println(array6)

	var array5 [20]int
	array5[0] = 3
	array5[1] = 4
	fmt.Println(array5)
	copy(array5[:], array6[:])
	fmt.Println(array5[:], array6[:])

}
