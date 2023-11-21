package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mbndr/figlet4go"

	"github.com/svarlamov/goyhfin"
)

func main() {
	uptime := time.NewTicker(5 * time.Minute)
	tickers := strings.Split(os.Args[1], ",")

	defer uptime.Stop()

	for {
		FetchData(tickers)
		<-uptime.C
	}

}

func FetchData(tickers []string) {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = "colossal"
	ascii.LoadFont("./assets")

	var data [][]string

	for _, ticker := range tickers {
		resp, err := goyhfin.GetTickerData(ticker, goyhfin.OneDay, goyhfin.ThirtyMinutes, false)
		if err != nil {
			fmt.Print("error fetching data:", err)
			continue
		}

		change := 100 - resp.Quotes[0].Open*100/resp.Quotes[len(resp.Quotes)-1].Close

		tickerString := string(strings.Trim(resp.Symbol, "^")[0]) + "  "
		valueString := strconv.FormatFloat(resp.Quotes[len(resp.Quotes)-1].Close, 'f', 2, 64) + "  "

		changeColor := "91"
		changeString := strconv.FormatFloat(change, 'f', 2, 64)
		if change >= 0 {
			changeColor = "92"
			changeString = "+" + changeString
		}

		renderTicker, _ := ascii.RenderOpts(tickerString, options)
		renderValue, _ := ascii.RenderOpts(valueString, options)
		renderChange, _ := ascii.RenderOpts(changeString, options)

		data = append(data, []string{renderTicker, renderValue, renderChange, changeColor})
	}
	
	PrintData(data)
}
