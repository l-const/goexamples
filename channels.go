package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	fun := func() {
		time.Sleep(time.Second * 10)
		fmt.Println("hello there")
		messages <- "ping"
	}

	go fun()
	msg := <-messages
	time.Sleep(time.Second * 3)
	fmt.Println(`there as well 
	from main, with recv msg: `, msg)
}
