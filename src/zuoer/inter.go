package zuoer

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Stringable interface {
	ToString() string
	// Size() int
}

func (c Country) ToString() string {
	return c.Name
}

func (c City) ToString() string {
	return c.Name
}

func PrintStr(s Stringable) {
	fmt.Println(s.ToString())
}

func Init() {
	co := Country{Name: "china"}
	ci := City{Name: "chongqing"}
	PrintStr(co)
	PrintStr(ci)

}
