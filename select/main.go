package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

var (
	NOMESSAGE_CHECK_INTERVAL = time.Second
	RETRY_INTERVAL           = 1 * time.Second
	closing                  = make(chan struct{})
)

func main() {
	timeoutTimer := time.NewTimer(NOMESSAGE_CHECK_INTERVAL)
	retryTimer := time.NewTimer(RETRY_INTERVAL)
	ch := make(chan int, 1)
	timecount := 0
	retycount := 0

	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	go func() {
		for {
			select {
			case i := <-ch:
				if i%10000 == 0 {
					fmt.Println(i)
				}
			case <-timeoutTimer.C:
				timecount++
				fmt.Println("timeout")

				timeoutTimer.Reset(NOMESSAGE_CHECK_INTERVAL)
			case <-retryTimer.C:
				retycount++
				fmt.Println("retry")
				retryTimer.Reset(RETRY_INTERVAL)
			case <-closing:
				return
			}
		}
	}()

	for i := 0; i < 100000000000; i++ {
		// time.Sleep(time.Microsecond * 500)
		ch <- i
		// time.Sleep(time.Minute * 10)
	}
}

func Stop() {
	closing <- struct{}{}
	<-closing
	close(closing)
}
