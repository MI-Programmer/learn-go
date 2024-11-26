package main

// import (
// 	"fmt"
// )

// func main() {
// 	// fruits := [...]string{"apple", "cherry", "watermelon", "orange", "purple"}

// 	// slice1 := fruits[2:5]
// 	// slice2 := fruits[:2]
// 	// slice3 := fruits[1:]
// 	// slice4 := fruits[:]

// 	// fmt.Println(slice1)
// 	// fmt.Println(slice2)
// 	// fmt.Println(slice3)
// 	// fmt.Println(slice4)

// 	days := [...]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

// 	daySlice1 := days[5:]
// 	fmt.Println(daySlice1)

// 	daySlice1[0] = "new friday"
// 	daySlice1[1] = "new saturday"
// 	fmt.Println(daySlice1)
// 	fmt.Println(days)

// 	daySlice2 := append(daySlice1, "new holiday")
// 	daySlice2[0] = "old friday"
// 	fmt.Println(daySlice1)
// 	fmt.Println(daySlice2)
// 	fmt.Println(days)

// 	var nums []int = make([]int, 2, 5)
// 	nums[0] = 1
// 	nums[1] = 2
// 	// nums[2] = 3 error

// 	fmt.Println(nums)
// 	fmt.Println(len(nums))
// 	fmt.Println(cap(nums))

// 	sliceNums1 := append(nums, 3)
// 	fmt.Println(sliceNums1)
// 	fmt.Println(len(sliceNums1))
// 	fmt.Println(cap(sliceNums1))

// 	sliceNums1[0] = 9
// 	fmt.Println(sliceNums1)
// 	fmt.Println(nums)

// 	fromSlice := days[:]
// 	toSlice := make([]string, len(fromSlice), cap(fromSlice))

// 	copy(toSlice, fromSlice)

// 	fmt.Println(fromSlice)
// 	fmt.Println(toSlice)

// 	thisArray := [...]int{1, 2, 3, 4, 5}
// 	thisSlice := []int{1, 2, 3, 4, 5}

// 	fmt.Println(thisArray)
// 	fmt.Println(thisSlice)
// }
