package ref

import (
	"fmt"
	"reflect"
)

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elem := v.Elem()
	n := elem.NumMethod()
	if n == 0 {
		fmt.Println("No methods")
		return
	}
	fmt.Println("Methods:")
	for i := 0; i < n; i++ {
		m := elem.Method(i)
		fmt.Printf("%s-", m.Name)
	}
	fmt.Printf("\n")
}
