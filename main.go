package main

import (
	"fmt"
	"time"
)

func goRoutine(quit chan string) {
	for {
		select {
		case <-quit:
			fmt.Println("stopped in goroutine")
			return
		default:
			fmt.Println("working, now wait for 1 second in goroutine")
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now())
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
	fmt.Println(fmt.Sprintf("main goroutine finished: %s", time.Now()))
}
