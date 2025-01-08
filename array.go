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

	// Slice
	slice := []int{1, 2, 3, 4, 5}
	println(slice[0], slice[1], slice[2], slice[3], slice[4])
	println("The length is %v with capacity %v\n",len(slice), cap(slice))
	slice = append(slice, 6)
	println("The length is %v with capacity %v",len(slice), cap(slice))
	println(slice)

	var createCapacity = make([]int, 5, 5)
	println("The length is %v with capacity %v",len(createCapacity), cap(createCapacity))
}