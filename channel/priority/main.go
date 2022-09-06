package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	done := make(chan struct{})
	notify1 := make(chan struct{})
	notify2 := make(chan struct{})
	notify3 := make(chan struct{})


	go func() {
		select {
		case <-done: //有一个完成就结束
			cancel()
		}
	}()

	//四个同时处理，优先级高的先处理完，后面的就不处理了



}

func handleOne(ctx context.Context, done, notify2 chan struct{}) {
	// do something
	if err := do(); err != nil {
		// 1失败了，通知2处理
		notify2 <- struct{}{}
		return
	}
	//成功后取消其他的协程
	done <- struct{}{}

}

func handleTwo(ctx context.Context, done, notify2, notify3 chan struct{}) {
	// do something
	select {
	case <-notify2:
		// do something
		err := do()
		if err != nil {
			// 2失败了，通知3处理
			notify3 <- struct{}{}
			return
		}
		//成功后取消其他的协程
		done <- struct{}{}
	case <-ctx.Done():
		return
	}
}

func handleThree() {

}

func handleFour() {

}

func do() error {
	fmt.Println("do something")
	return nil
}
