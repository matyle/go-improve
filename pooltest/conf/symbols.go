package conf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"test/pooltest/utils"
	"time"

	"cryptobroker/feemanager/common"

	"github.com/sirupsen/logrus"
)

type SymbolConf struct {
	Symbol            string `json:"symbol"`
	BasePrecision     int    `json:"basePrecision"`
	QuotePrecision    int    `json:"quotePrecision"`
	SpendingPrecision int    `json:"spendingPrecision"`
	MinAmount         string `json:"minAmount"`
	MinQty            string `json:"minQty"`
	MaxQty            string `json:"maxQty"`
	MinSpending       string `json:"minSpending"`
	MaxSpending       string `json:"maxSpending"`
	MinDumping        string `json:"minDumping"`
	MaxDumping        string `json:"maxDumping"`
	MaxPrice          string `json:"maxPrice"`
	BuyCeiling        string `json:"buyCeiling"`
	SellFloor         string `json:"sellFloor"`
	SpendingInterval  string `json:"spendingInterval"`
	DumpingInterval   string `json:"dumpingInterval"`
	MakerFeeRate      string `json:"makerFeeRate"`
	TakerFeeRate      string `json:"takerFeeRate"`

	PriceShift int `json:"matcherPriceShifting"`
	QtyShift   int `json:"matcherQtyShifting"`

	PriceFactor int `json:"priceMultiplier,omitempty"`
	QtyFactor   int `json:"qtyMultiplier,omitempty"`

	MatcherVersion int `json:"matcherVersion"`
}

type SymbolsConfFile struct {
	Symbols []SymbolConf `json:"symbols"`
}

var (
	MapInterval = map[string]int64{
		"1M":  60,      // 1 min
		"5M":  300,     // 5 min
		"15M": 900,     // 15 min
		"30M": 1800,    // 30 min
		"60M": 3600,    // 1 hour
		"4H":  14400,   // 4 hour
		"8H":  28800,   // 8 hour
		"12H": 43200,   // 12 hour
		"1D":  86400,   // 1 day
		"1W":  604800,  // 1 week
		"1m":  2592000, // 1 month
	}

	bufCap = 3 >> 20 // 3M

	DepthLimit = []int{0, 5, 10, 20, 50, 100, 150, 500, 1000}

	SymbolsBuf           bytes.Buffer
	BytesBufPool         = utils.NewBytesBufferPool(bufCap)
	symbolsConf          SymbolsConfFile
	MapSymbolsConf       = make(map[string]*SymbolConf)
	SymbolList           []string
	subscribeSymbolsChan []chan struct{}
)

func GetAllSymbolsConf() map[string]*SymbolConf {
	maps := make(map[string]*SymbolConf, 0)
	for _, v := range symbolsConf.Symbols {
		vv := v
		maps[v.Symbol] = &vv
	}
	return maps
}

func LoadSymbols() {
	funcLoad := func() {
		tStart := time.Now()
		content, err := ReadFileOrGzipFile("./symbolconf.json") // k8s configmap配置文件
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(content, &symbolsConf)
		if err != nil {
			panic(err)
		}

		// 统计version==0的symbol数量
		versionCount := 0
		for _, v := range symbolsConf.Symbols {
			if v.MatcherVersion == 0 {
				versionCount++
			}
		}

		BytesBuf := BytesBufPool.Get()
		BytesBuf.Write(content)
		defer BytesBufPool.Put(BytesBuf)

		logrus.Printf("load symbols conf cost: %v, versionCount: %v", time.Since(tStart), versionCount)
		logrus.Printf("symbols conf: %v", BytesBuf)
	}
	funcLoad()

	go func() {
		for {
			select {
			case <-time.After(time.Second * 15):
				funcLoad()
			}
		}
	}()
}

func RegisterSubChan(ch chan struct{}) {
	subscribeSymbolsChan = append(subscribeSymbolsChan, ch)
}

func ReadFileOrGzipFile(file string) ([]byte, error) {
	b, err := ioutil.ReadFile(file + ".gz")
	if err != nil {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
	return common.UnGzipData(b)
}
