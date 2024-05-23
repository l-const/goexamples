package main

import "fmt"

func init() {
	fmt.Println("init runs first")
}

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// but then this block right?
	fmt.Println(<-messages)
	// actually since we block main routine
	// we get a deadlock with message
	// fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [chan receive]:
	// main.main()
	//    /Users/kostas/dev/gobyexamples/channels_buf.go:16 +0xf8
	// exit status 2

}
