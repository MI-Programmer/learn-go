package helper

import "fmt"

func HelloWorld(name string) string {
	return "Hello " + name
}

func FullName(firstName, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func Sum(num1, num2 int) int {
	return num1 + num2
}
