// Package main provides ...
package main

func main() {
	// zuoer.Init()
	// zuoer.InitOptions()
	// zuoer.Echo([]int{1, 2, 3, 4, 5})
	// nums := []int{1, 2, 3, 4, 5}
	//非 pipeline 方法

	//1+3+5
	// for v := range zuoer.Sum(zuoer.Odd(zuoer.Echo(nums))) {
	// 	println(v)
	// }

	//使用 pipeline 方法

	//
	ch := make(chan int)
	// go func() {
	// 	ch <- ch3.NumToChannel()
	// }()
	ch <- 1
	println(<-ch)

}
