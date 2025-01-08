package main

import (
	"fmt"
)

func main() {
	// Array
	var arr [5]int
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	arr[3] = 400
	arr[4] = 500
	println(arr[0], arr[1], arr[2], arr[3], arr[4])

	// Array literal
	arr2 := [5]int{1, 2, 3, 4, 5}
	println(arr2[0], arr2[1], arr2[2], arr2[3], arr2[4])

	// Array length
	println(len(arr2))

	// Array of arrays
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	println(arr3[0][0], arr3[0][1], arr3[1][0], arr3[1][1])

	// Slice
	slice := []int{1, 2, 3, 4, 5}
	println(slice[0], slice[1], slice[2], slice[3], slice[4])
	fmt.Printf("The length is %v with capacity %v\n", len(slice), cap(slice))
	slice = append(slice, 6)
	fmt.Printf("The length is %v with capacity %v", len(slice), cap(slice))
	println(slice)

	var createCapacity = make([]int, 5, 5)
	fmt.Printf("The length is %v with capacity %v", len(createCapacity), cap(createCapacity))

	// maps
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	println(m["one"], m["two"], m["three"], m["four"]) // get 0 if key not found

	//delete key
	// delete(m, "one")
	// println(m["one"], m["two"], m["three"], m["four"]) // get 0 if key not found

	// loops
	for name, value := range m { // NOTE: order is not guaranteed in maps
		fmt.Printf("Name: %v Value: %v \n", name, value)
	}

	// for loop
	for i := 0; i < 5; i++ {
		println(i)
	}
}
