package main

import (
	"fmt"
	"time"

	"github.com/svarlamov/goyhfin"
)

func (oculatus *Oculatus) FetchData() {
	uptime := time.NewTicker(5 * time.Minute)
	defer uptime.Stop()

	for ; true; <-uptime.C {
		oculatus.mutex.Lock()

		for i, ticker := range oculatus.Tickers {
			resp, err := goyhfin.GetTickerData(ticker.Symbol, goyhfin.OneDay, goyhfin.ThirtyMinutes, false)

			if err != nil {
				fmt.Print("error fetching data:", err)
				continue
			}

			tickerData := &TickerData{
				Symbol: ticker.Symbol,
				Price:  resp.Quotes[len(resp.Quotes)-1].Close,
				Change: 100 - resp.Quotes[0].Open*100/resp.Quotes[len(resp.Quotes)-1].Close,
			}

			oculatus.Tickers[i] = *tickerData
		}

		oculatus.mutex.Unlock()
		oculatus.trigger <- true
	}
}
