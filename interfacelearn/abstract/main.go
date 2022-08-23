package main

import "fmt"

type Source interface {
	Read() interface{}
}

type Index struct {
	symbol string
	price  int
}

type Mark struct {
	symbol    string
	makeprice float64
}

type SourceCache struct {
	source []Source
}

func (sc *Index) Read() interface{} {
	return sc
}

func (m *Mark) Read() interface{} {
	return m
}

func Gets(sc SourceCache) {
	for _, source := range sc.source {
		fmt.Printf("%v\n", source.Read())
	}
}

func Get(s Source) {
	fmt.Printf("%v\n", s.Read())
}

func main() {
	indexes := []*Index{
		{"GOOG", 500},
		{"MSFT", 50},
		{"FB", 150},
		{"AAPL", 200},
	}
	marks := []*Mark{
		{"GOOG", 500.23},
		{"MSFT", 50.2},
		{"FB", 150.1},
		{"AAPL", 200.3},
	}
	sc := SourceCache{}
	msc := SourceCache{}

	for _, index := range indexes {
		sc.source = append(sc.source, index)
		Get(index)
	}

	for _, mark := range marks {
		msc.source = append(msc.source, mark)
		Get(mark)
	}

	// Gets(sc)
	// Gets(msc)

}
