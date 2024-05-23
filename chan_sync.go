package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool, 1)
	go worker(done)
	fmt.Println("main: started waiting")
	// arewe := <-done
	// fmt.Println("worker received: ", arewe)

}

// a channel whose content is boolean
func worker(done chan bool) {
	fmt.Print("working ...")
	time.Sleep(time.Second * 2)
	fmt.Println("worker: done")
	done <- true
}
