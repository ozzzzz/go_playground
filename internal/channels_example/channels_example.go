package main

import (
	"fmt"
	"time"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	go produceEven(even)
	go produceOdd(odd)
	go produceQuit(quit)
	for {
		select {
		case e := <-even:
			fmt.Println("Even:", e)
		case o := <-odd:
			fmt.Println("Odd:", o)
		case <-quit:
			fmt.Println("Quitting")
			return
		default:
			fmt.Println("Waiting")
		}
	}
}

func produceEven(nums chan int) {
	for i := 0; ; i += 2 {
		nums <- i
	}
}

func produceOdd(nums chan int) {
	for i := 1; ; i += 2 {
		nums <- i
	}
}

func produceQuit(quit chan bool) {
	time.Sleep(5 * time.Second)
	quit <- true
}
