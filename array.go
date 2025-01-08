package main

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
}