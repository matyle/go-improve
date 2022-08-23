package main

import (
	"strconv"
)

type Test struct {
	Name string
	Age  int
}

func (t Test) String() string {
	return "Name: " + t.Name + " Age: " + strconv.FormatInt(int64(t.Age), 10)
}

func main() {

	test := &Test{
		Name: "test",
		Age:  10,
	}
	println(test)
	println(test.String())
}
