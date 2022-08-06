package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(j int) {
	println("work", j)
	time.Sleep(time.Second * time.Duration(j))
}

// WaitAndExit waits for a worker to finish and then exits.
func spawn(f func(int)) chan string {
	quit := make(chan string)
	// Start a goroutine to do the work.
	go func() {
		var job chan int // job channel
		for {
			select {
			case j := <-job:
				f(j)
			case <-quit:
				quit <- "ok"
			}
		}
	}()
	return quit
}
func spawnGroup(n int, f func(int)) chan struct{} {
	quit := make(chan struct{})
	job := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprintf("worker %d", i)
			wg.Add(1)
			for {
				if _, ok := <-job; !ok {
					println("worker-done", name)
					return
				}
				f(i)
			}
		}(i)
	}

	// Start a goroutine to wait
	go func() {
		<-quit //等待 1
		close(job)
		wg.Wait()
		quit <- struct{}{}
	}()

	return quit
}

func main() {
	// quit := spawn(worker)
	quit := spawnGroup(5, worker)

	println("start")
	time.Sleep(time.Second * 5)
	println("notify and exit")
	// quit <- "exit"
	quit <- struct{}{} //1
	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	select {
	case <-timer.C:
		println("timeout")
	case <-quit:
		println("work done ok")
	}
}
