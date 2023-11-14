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
		PrintTickerData(tickers)
		<-uptime.C
	}

}

func PrintTickerData(tickers []string) {
	//print this gibberish to clear terminal
	fmt.Print("\033[H\033[2J")

	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = "univers"
	ascii.LoadFont("./assets")

	for _, ticker := range tickers {
		resp, err := goyhfin.GetTickerData(ticker, goyhfin.OneDay, goyhfin.ThirtyMinutes, false)
		if err != nil {
			fmt.Println("error fetching data:", err)
			panic(err)
		}

		tickerString := string(strings.Trim(ticker, "^")[0]) + ": "
		valueString := strconv.FormatFloat(resp.Quotes[0].High, 'f', 2, 64)

		renderTicker, _ := ascii.RenderOpts(tickerString, options)
		renderValue, _ := ascii.RenderOpts(valueString, options)

		fmt.Printf("%4s, %10s", renderTicker, renderValue)
	}
}
