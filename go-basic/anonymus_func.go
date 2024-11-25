package main

// type Blacklist func(string) bool

// func register(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("Youre blocked", name)
// 	} else {
// 		fmt.Println("Youre welcome", name)
// 	}
// // }

// type Callback func(int) bool

// func find(values []int, callback Callback) (int, int) {
// 	for i, value := range values {
// 		if callback(value) {
// 			return value, i
// 		}
// 	}

// 	return -1, -1
// }

// func main() {
// blacklist := func(name string) bool {
// 	return name == "anjing"
// }

// register("ilham", blacklist)
// register("anjing", func(name string) bool {
// 	return name == "anjing"
// })

// nums := []int{1, 2, 3, 4, 5}

// value, index := find(nums, func(num int) bool { return num == 10 })

// fmt.Println(value, index)
// }
