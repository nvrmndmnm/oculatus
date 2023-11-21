package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mbndr/figlet4go"
	"github.com/olekukonko/tablewriter"
)

func RenderFigletStrings(oculatus *Oculatus) {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = "colossal"
	ascii.LoadFont("./assets")

	for range oculatus.trigger {
		var data [][]string
		oculatus.mutex.RLock()

		for _, tickerData := range oculatus.Tickers {
			symbolString := string(strings.Trim(tickerData.Symbol, "^")[0]) + "  "
			priceString := strconv.FormatFloat(tickerData.Price, 'f', 2, 64) + "  "
			changeString := strconv.FormatFloat(tickerData.Change, 'f', 2, 64) + "  "
			changeColor := tablewriter.FgHiRedColor

			if tickerData.Change >= 0 {
				changeColor = tablewriter.FgHiGreenColor
				changeString = "+" + changeString
			}

			figletSymbol, _ := ascii.RenderOpts(symbolString, options)
			figletPrice, _ := ascii.RenderOpts(priceString, options)
			figletChange, _ := ascii.RenderOpts(changeString, options)

			data = append(data, []string{figletSymbol, figletPrice, figletChange, fmt.Sprint(changeColor)})
		}

		oculatus.mutex.RUnlock()

		printTable(data)
	}
}

func printTable(data [][]string) {
	//print this gibberish to clear terminal
	fmt.Print("\033[H\033[2J\n")

	tableString := &strings.Builder{}

	table := *tablewriter.NewWriter(tableString)

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)
	table.SetBorder(false)
	table.SetColumnSeparator("")	

	for _, v := range data {
		colorCode, _ := strconv.Atoi(v[len(v)-1])

		table.Rich(v[:len(v)-1], []tablewriter.Colors{
			{tablewriter.Bold, colorCode},
			{tablewriter.Bold, colorCode},
			{tablewriter.Bold, colorCode},
		})
	}

	table.Render()
	fmt.Println(tableString.String())
}
