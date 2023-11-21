package main

import (
	"os"
	"strings"
	"sync"
)

type TickerData struct {
	Symbol string
	Price  float64
	Change float64
}

type Oculatus struct {
	Tickers []TickerData
	trigger chan bool
	mutex   sync.RWMutex
}

func main() {
	inputSymbols := strings.Split(os.Args[1], ",")
	oculatus := Oculatus{
		trigger: make(chan bool, 1),
	}

	for _, symbol := range inputSymbols {
		tickerData := &TickerData{
			Symbol: symbol,
		}
		oculatus.Tickers = append(oculatus.Tickers, *tickerData)
	}

	go oculatus.FetchData()
	go RenderFigletStrings(&oculatus)

	select {}
}
