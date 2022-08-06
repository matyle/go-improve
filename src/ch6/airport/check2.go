package airport

import "time"

func IdCheck2(id int) int {
	print("goroutine-", id)
	time.Sleep(time.Millisecond * time.Duration(idCheckCost))
	println("\tId check done")
	return idCheckCost
}

func BodyCheck2() int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckCost))
	println("\tBody check done")
	return bodyCheckCost
}

func XRayCheck2() int {
	time.Sleep(time.Millisecond * time.Duration(XRayCheckCost))
	println("\tX-Ray check done")
	return XRayCheckCost
}

func AirportsCheck2() int {
	println("start checking airports security...")
	total := 0
	total += IdCheck()
	total += BodyCheck()
	total += XRayCheck()
	println("all check done")
	return total
}
