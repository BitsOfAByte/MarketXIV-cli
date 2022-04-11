package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BitsOfAByte/marketxiv/cli"
)

func main() {

	if len(os.Args) < 2 {
		cli.ShowQuickHelp()
		return
	}

	switch os.Args[1] {
	case "--help":
		cli.ShowHelpPage()
	case "--info":
		cli.ShowInfo()
	case "--update":
		fmt.Println("Not implemented yet.")
	case "item":
		ItemLookup()
	case "tax":
		MarketTax()
	default:
		cli.ShowQuickHelp()
	}
}

func ItemLookup() {
	if len(os.Args) < 4 {
		fmt.Println("marketxiv: specify a server and item name. Use --help for more information.")
		return
	}

	item := cli.MarketItem{ItemName: strings.Join(os.Args[3:], " "), Server: os.Args[2]}
	item.ShowListings(10)
}

func MarketTax() {
	if len(os.Args) < 3 {
		fmt.Println("marketxiv: specify a server. Use --help for more information.")
		return
	}

	market := cli.Market{Server: os.Args[2]}
	market.TaxRates()
}
