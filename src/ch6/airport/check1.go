package airport

import "time"

const (
	idCheckCost   = 60
	bodyCheckCost = 120
	XRayCheckCost = 180
)

func IdCheck() int {
	time.Sleep(time.Millisecond * time.Duration(idCheckCost))
	println("\tId check done")
	return idCheckCost
}

func BodyCheck() int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckCost))
	println("\tBody check done")
	return bodyCheckCost
}

func XRayCheck() int {
	time.Sleep(time.Millisecond * time.Duration(XRayCheckCost))
	println("\tX-Ray check done")
	return XRayCheckCost
}

func AirportsCheck() int {
	println("start checking airports security...")
	total := 0
	total += IdCheck()
	total += BodyCheck()
	total += XRayCheck()
	println("all check done")
	return total
}

// func main() {
// 	total := 0
// 	passage := 30
// 	for i := 0; i < passage; i++ {
// 		total += AirportsCheck()
// 	}
// 	println("total cost:", total)
// }
