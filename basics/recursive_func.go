package main

// func factorialLoop(num int) int {
// 	result := 1

// 	for i := num; i > 0; i-- {
// 		result *= i
// 	}

// 	return result
// }

// func factorialRecursive(num int) int {
// 	if num == 1 {
// 		return 1
// 	} else {
// 		return num * factorialRecursive(num-1)
// 	}
// }

// func isPalindrome(value string) bool {
// 	length := len(value)
// 	if length <= 1 {
// 		return true
// 	}
// 	if value[0] != value[length-1] {
// 		return false
// 	}
// 	return isPalindrome(value[1 : length-1])
// }

// func isPalindrome(value string) bool {
// 	length := len(value)
// 	if length <= 1 {
// 		return true
// 	}
// 	return value[0] == value[length-1] && isPalindrome(value[1:length-1])
// }

// func main() {
// 	fmt.Println(factorialLoop(3))
// 	fmt.Println(factorialRecursive(3))
// 	fmt.Println(isPalindrome("agasga"))
// 	fmt.Println(isPalindrome("aasssasssaa"))
// }
