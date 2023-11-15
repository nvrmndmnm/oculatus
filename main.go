package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mbndr/figlet4go"
	"github.com/olekukonko/tablewriter"
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
	options.FontName = "univers"
	ascii.LoadFont("./assets")

	var data [][]string

	for _, ticker := range tickers {
		resp, err := goyhfin.GetTickerData(ticker, goyhfin.OneDay, goyhfin.ThirtyMinutes, false)
		if err != nil {
			fmt.Print("error fetching data:", err)
			continue
		}

		tickerString := string(strings.Trim(ticker, "^")[0]) + "  "
		valueString := strconv.FormatFloat(resp.Quotes[0].High, 'f', 2, 64)

		renderTicker, _ := ascii.RenderOpts(tickerString, options)
		renderValue, _ := ascii.RenderOpts(valueString, options)

		data = append(data, []string{renderTicker, renderValue})
	}

	printData(data)
}

func printData(data [][]string) {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)

	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnSeparator("")
	table.SetBorder(false)
	table.SetTablePadding("\t")

	table.AppendBulk(data)
	table.Render()

	//print this gibberish to clear terminal
	fmt.Print("\033[H\033[2J")

	fmt.Println(tableString.String())
}
