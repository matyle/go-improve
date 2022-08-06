// for range
package ch3

import (
	"fmt"
	"time"
)

func SliceChange() {
	s := []int{1, 2, 3, 4, 5}
	r := make([]int, len(s))
	for i, v := range s {
		if i == 0 {
			s = append(s, 6)
			s = append(s, 7)
		}
		r[i] = v
	}
	fmt.Println("r= ", r)
	fmt.Println("s= ", s)

}

func StringRange() {
	s := "hello"
	for i, v := range s {
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}
}

func BrakeTest() {
	exit := make(chan interface{})

	go func() {
	loop:
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exit")
				break loop
			}

		}
	}()
	time.Sleep(3 * time.Second)
	exit <- struct{}{}
	time.Sleep(3 * time.Second)
}
