package goroutines_test

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// channel := make(chan string)
	// defer close(channel)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	channel <- "ilham ganteng"
	// 	fmt.Println("Done send data to channel")
	// }()

	// data := <-channel
	// fmt.Println(data)

	// time.Sleep(5 * time.Second)

	channel := make(chan []string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- []string{"ilham", "ganteng"}
		fmt.Println("Done send data to channel 1")
		channel <- []string{"hello", "world"}
		fmt.Println("Done send data to channel 2")
	}()

	data1 := <-channel
	data2 := <-channel
	fmt.Println(data1)

	for _, v := range data2 {
		fmt.Println(v)
	}

	time.Sleep(5 * time.Second)
}
