package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func PrintData(data [][]string) {
	//print this gibberish to clear terminal
	fmt.Print("\033[H\033[2J")

	tableString := &strings.Builder{}

	table := *tablewriter.NewWriter(tableString)

	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnSeparator("")
	table.SetBorder(false)

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
