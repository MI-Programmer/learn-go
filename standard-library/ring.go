package main

// import (
// 	"container/ring"
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var data *ring.Ring = ring.New(5)

// 	for i := 1; i <= data.Len(); i++ {
// 		data.Value = "Value " + strconv.Itoa(i)
// 		data = data.Next()
// 	}

// 	// data.Value = "Value 1"
// 	// data = data.Next()

// 	// data.Value = "Value 2"
// 	// data = data.Next()

// 	// data.Value = "Value 3"
// 	// data = data.Next()

// 	// data.Value = "Value 4"
// 	// data = data.Next()

// 	// data.Value = "Value 5"

// 	data.Do(func(e any) {
// 		fmt.Println(e)
// 	})
// }

// func main() {
// 	data := ring.New(6)

// 	for i := 1; i <= data.Len(); i++ {
// 		data.Value = strconv.Itoa(i) + ". data default"
// 		data = data.Next()
// 	}

// 	data.Do(func(a any) {
// 		fmt.Println(a)
// 	})
// }
