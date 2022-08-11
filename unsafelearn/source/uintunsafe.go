package source

import (
	"fmt"
	"runtime"
)

type Foo struct {
	name string
}

func finalizer(f *Foo) {
	fmt.Printf("Foo:[%s] is finalized\n", f.name)
}

func NewFoo(name string) *Foo {
	var f = Foo{name: name}
	runtime.SetFinalizer(&f, finalizer)
	return &f
}

func Allocate() *[1000000]uint64 {
	var a = [1000000]uint64{}
	return &a
}
