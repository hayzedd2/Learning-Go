package main

import (
	"fmt"
	"time"
)

func main() {
	// chanStatuses := make([]chan bool, 4)
	chanStatus := make(chan bool)
	go greet("Hello there", chanStatus)
	// chanStatuses[1] = make(chan bool)
	go greet("Hi there", chanStatus)
	// chanStatuses[2] = make(chan bool)
	go slowGreet("This is a slow greeting", chanStatus)
	// chanStatuses[3] = make(chan bool)
	go greet("Good afternoon", chanStatus)

	// for _, chanStatus := range chanStatuses {
	// 	<-chanStatus
	// }
	for range chanStatus{

	}
}

func slowGreet(msg string, chanStatus chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(msg)
	chanStatus <- true
	close(chanStatus)
}
func greet(msg string, chanStatus chan bool) {
	startTime := time.Now()
	elapsedTime := time.Since(startTime)
	fmt.Println(msg, "Elapsed time:", elapsedTime)
	chanStatus <- true

}
