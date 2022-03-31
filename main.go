package main

import (
	"fmt"
	"time"
)

func goRoutine(ch chan string) {
	for {
		fmt.Println("working, now wait for 1 second in goroutine")
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now())
		data := <-ch
		fmt.Println(fmt.Sprintf("signal from main goroutine: %s", data))
		if data != "stop" {
			fmt.Println("working in goroutine")
		} else {
			fmt.Println("stopped in goroutine")
			return
		}
	}
}

func main() {
	fmt.Println(fmt.Sprintf("main goroutine starting: %s", time.Now()))
	ch := make(chan string)
	go goRoutine(ch)
	fmt.Println("wait for 20 second in main goroutine")
	time.Sleep(20 * time.Second)
	ch <- "stop"
	time.Sleep(1 * time.Second)
	fmt.Println(fmt.Sprintf("main goroutine finished: %s", time.Now()))
}
