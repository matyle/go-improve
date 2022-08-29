package main

import (
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
	sync.Mutex
}

type Students struct {
	ps []*Person
	sync.Mutex
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	// defer println("defer", i)
	// 	err
	// 	defer func() {
	// 		println("defer", i)
	// 	}()
	// }
	go func() {
		err := 0
		defer func() {
			println("defer", err)
		}()
		err = 10
	}()
	time.Sleep(time.Second * 10)
}

func testfor(i int) {
	println("testfor", i)
	return
}
