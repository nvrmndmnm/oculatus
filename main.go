package main

import (
	"fmt"
	"os"
	"strings"
	"time"

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

	fmt.Printf("+%s+\n", strings.Repeat("-", 50))

	for index, ticker := range tickers {
		resp, err := goyhfin.GetTickerData(ticker, goyhfin.OneDay, goyhfin.FiveMinutes, false)
		if err != nil {
			fmt.Println("error fetching data:", err)
			panic(err)
		}

		fmt.Printf("| %3c %10.2f ", strings.Trim(ticker, "^")[0], resp.Quotes[0].High)
		if (index+1)%3 == 0 {
			fmt.Printf("|\n+%s+\n", strings.Repeat("-", 50))
		}
	}
}
