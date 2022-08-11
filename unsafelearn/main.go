package main

import (
	"fmt"
	"goimpove/go-improve/unsafelearn/source"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	p1 := uintptr(unsafe.Pointer(source.NewFoo("FooRefByUintPtr")))
	p2 := unsafe.Pointer(source.NewFoo("FooRefByUnsafePointer"))

	for i := 0; i < 10; i++ {
		source.Allocate()
		q1 := (*source.Foo)(unsafe.Pointer(p1)) //uintptr 转换
		fmt.Printf("object ref by uintptr: %+v\n", *q1)
		q2 := (*source.Foo)(p2)
		fmt.Printf("object ref by pointer: %+v\n", *q2)

		runtime.GC()
		time.Sleep(time.Second)

	}
}
