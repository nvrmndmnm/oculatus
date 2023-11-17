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

	printData(data)
}

func printData(data [][]string) {
	//print this gibberish to clear terminal
	fmt.Print("\033[H\033[2J")

	for _, v := range data {
		tableString := &strings.Builder{}
		table := tablewriter.NewWriter(tableString)

		table.SetHeader([]string{"", "", ""})
		table.SetHeaderLine(false)
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetColumnSeparator("")
		table.SetBorder(false)
		table.SetTablePadding("\t")

		colorCode, _ := strconv.Atoi(v[len(v)-1])

		table.SetColumnColor(
			tablewriter.Colors{tablewriter.Bold, colorCode},
			tablewriter.Colors{tablewriter.Bold, colorCode},
			tablewriter.Colors{tablewriter.Bold, colorCode},
		)

		table.Append(v[:len(v)-1])
		table.Render()
		fmt.Println(tableString.String())
	}
}
