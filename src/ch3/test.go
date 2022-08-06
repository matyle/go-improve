package ch3

import "go-improve/src/ch3/ref"

type Inter interface {
	M1()
	M2()
}

type T struct {
	name string
}

func (t *T) M1() {
	println(t.name)
}
func (t *T) M2() {
	println(t.name)
}

func Dump() {
	var e interface{} = &T{"Hello, world!"}
	var e2 interface{} = &T{"nice"}
	var i Inter = &T{"slq"}
	// ref.DumpMethodSet(e)
	ref.DumpMethodSet(&i) //为什么要取地址？
	println("e", e)
	println("e2", e2)
	println("i", i)
	println("e==i", e == i)
	i.M2()
	i.M1()

}
