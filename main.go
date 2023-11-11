package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/svarlamov/goyhfin"
)

func main() {
	var tickers = strings.Split(os.Args[1], ",")
	fmt.Printf("+%s+|\n", strings.Repeat("-", 50))
	for index, ticker := range tickers {
		resp, err := goyhfin.GetTickerData(ticker, goyhfin.OneDay, goyhfin.FiveMinutes, false)
		if err != nil {
			fmt.Println("Error fetching Yahoo Finance data:", err)
			panic(err)
		}

		fmt.Printf("| %3c %10.2f ", strings.Trim(ticker, "^")[0], resp.Quotes[0].High)
		if (index+1)%3 == 0 {
			fmt.Printf("\n+%s+|\n", strings.Repeat("-", 50))
		}
	}
}
