package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(TableName())
}

func TableName() string {
	//月初1号需要查找到上个月的表
	now := time.Now()
	tm := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.UTC).
		AddDate(0, 0, -1)
	logrus.WithFields(logrus.Fields{
		"year":  tm.Year(),
		"month": tm.Month(),
	}).Debug("TableName")
	return fmt.Sprintf("binance_trades_%s", tm.Format("200601"))
	// return fmt.Sprintf("binance_trades_%s", "202101")
}
