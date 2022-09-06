package main

import "fmt"

func main() {
	var removePos []int
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	delnum := 4
	var nRepeat int

	for i, k := range nums {
		if k > delnum {
			nRepeat++
			removePos = append(removePos, i)
		} else {
		}
	}

	for i, pos := range removePos {
		fmt.Println(i, pos)
		nums = append(nums[:pos-i], nums[pos-i+1:]...)
	}
	for _, num := range nums {
		fmt.Println("num:", num)
	}

	return

}
