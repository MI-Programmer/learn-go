package main

// func divide(num int, divider int) (int, error) {
// 	if divider == 0 {
// 		return 0, errors.New("Dividing with zero!")
// 	} else {
// 		return num / divider, nil
// 	}
// }

// func main() {
// 	result, err := divide(10, 0)

// 	if err == nil {
// 		fmt.Println("Result", result)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}
// }

// func filter(value string) (string, error) {
// 	if value == "" {
// 		return "", errors.New("String value is empty!")
// 	}

// 	if value == "anjing" {
// 		return "...", nil
// 	} else {
// 		return value, nil
// 	}
// }

// func main() {
// 	result, err := filter("")

// 	if err == nil {
// 		fmt.Println(result)
// 	} else {
// 		fmt.Println("Error:", err)
// 	}
// }

// type Student struct {
// 	Name, Age string
// }

// func newStudent(name, age string) (*Student, error) {
// 	if name == "" || age == "" {
// 		return nil, errors.New("Name or Age is not provided")
// 	}

// 	return &Student{name, age}, nil
// }

// func main() {
// 	result1, _ := newStudent("ilham", "20")
// 	result2, err := newStudent("ilham", "")

// 	fmt.Println(result1)
// 	fmt.Println(result2, err)
// }
