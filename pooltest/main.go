package main

import (
	"test/pooltest/conf"
	"time"
)

func main() {
	conf.LoadSymbols()
	time.Sleep(1 * time.Minute)
}
