package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var now time.Time = time.Now()
// 	fmt.Println(now)

// 	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
// 	fmt.Println(utc)
// 	fmt.Println(utc.Local())

// 	formatter := "2006-01-02 15:04:05"

// 	value := "2020-10-10 10:10:10"
// 	// value := "ASAL"
// 	valueTime, err := time.Parse(formatter, value)
// 	if err == nil {
// 		fmt.Println(valueTime)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}

// 	fmt.Println(valueTime.Minute())
// 	fmt.Println(valueTime.Hour())
// 	fmt.Println(valueTime.Day())
// 	fmt.Println(valueTime.Month())
// 	fmt.Println(valueTime.Year())
// }

// func main() {
// 	var now time.Time = time.Now()
// 	fmt.Println(now)
// 	fmt.Println(now.Local())
// 	fmt.Println(now.Year())

// 	var utc time.Time = time.Date(2024, time.January, 12, 10, 10, 10, 0, time.Local)
// 	fmt.Println(utc)
// 	fmt.Println(utc.String())

// 	formatter := "2006-01-02 15:04:05"
// 	value := "2023-09-29 20:22:59"

// 	valueTime, err := time.Parse(formatter, value)
// 	if err != nil {
// 		fmt.Println("Error", err.Error())
// 	} else {
// 		fmt.Println(valueTime)
// 		fmt.Println(valueTime.Month())
// 	}
// }
