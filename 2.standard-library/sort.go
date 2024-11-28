package main

// import (
// 	"fmt"
// 	"sort"
// )

// type User struct {
// 	Name string
// 	Age  int
// }

// type UserSlice []User

// func (u UserSlice) Len() int {
// 	return len(u)
// }

// func (u UserSlice) Less(i, j int) bool {
// 	return u[i].Age < u[j].Age
// }

// func (u UserSlice) Swap(i, j int) {
// 	u[i], u[j] = u[j], u[i]
// }

// func main() {
// 	users := UserSlice{
// 		{"ilham", 20},
// 		{"mi", 10},
// 		{"programmer", 40},
// 		{"mi ganteng", 5},
// 	}

// 	sort.Sort(users)

// 	fmt.Println(users)
// }

// type Student struct {
// 	Class string
// 	Score int
// }

// type StudentSlice []Student

// func (s StudentSlice) Len() int {
// 	return len(s)
// }

// func (s StudentSlice) Less(i, j int) bool {
// 	return s[i].Score < s[j].Score
// }

// func (s StudentSlice) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// func main() {
// 	students := []Student{
// 		{"A", 30},
// 		{"A", 90},
// 		{"A", 10},
// 		{"D", 40},
// 		{"C", 20},
// 	}

// 	sort.Sort(StudentSlice(students))

// 	fmt.Println(students)
// }
