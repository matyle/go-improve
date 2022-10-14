package main

import "fmt"

var testChan = make(chan struct{})

func main() {
	go mainLoop()

	testChan <- struct{}{}
	testChan <- struct{}{}
	testChan <- struct{}{}
	<-testChan
	close(testChan)
	fmt.Println("non blocked")
}

func mainLoop() {
	count := 0
	for range testChan {
		fmt.Println("receive")
		count++
		if count >= 3 {
			break
		}
	}
	// testChan <- struct{}{}
}
