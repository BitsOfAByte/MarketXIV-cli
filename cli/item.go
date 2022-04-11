package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/olekukonko/tablewriter"
)

// Information about an item that is marketable.
type MarketItem struct {
	ItemName string
	Server   string
}

// Get information about the item from the market.
func (i MarketItem) ShowListings(limit int) {

	// Fetch the data for the item
	itemApiResponse := backend.FetchItem(i.ItemName)

	if len(itemApiResponse.Results) == 0 {
		fmt.Printf("marketxiv: No listings found for %s.", i.ItemName)
		return
	}

	itemData := itemApiResponse.Results[0]
	marketItemData := backend.FetchMarketItem(i.Server, itemData.ID, limit)

	if len(marketItemData.Listings) == 0 {
		fmt.Printf("marketxiv: %s is not a marketable item.", i.ItemName)
		return
	}

	// Prepare a table to display the data
	listingData := make([][]string, len(marketItemData.Listings))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Quality", "Price", "Quantity", "Total", "Retainer", "World"})

	// Loop through the listings and add them their data to the table
	for _, listing := range marketItemData.Listings {

		// Handle the quality boolean and convert it to a string.
		var quality string
		switch listing.Hq {
		case true:
			quality = "HQ"
		default:
			quality = "NQ"
		}

		// When there is no world name, assume the world is only the server entered.
		if listing.WorldName == "" {
			listing.WorldName = strings.ToTitle(strings.ToLower(i.Server))
		}

		// Add the listing to the table
		listingData = append(listingData, []string{
			quality,
			strconv.Itoa(listing.PricePerUnit),
			strconv.Itoa(listing.Quantity),
			strconv.Itoa(listing.Total),
			listing.RetainerName,
			listing.WorldName,
		})
	}

	// Show the table
	table.AppendBulk(listingData)
	fmt.Printf("Showing listings for %s\n", i.ItemName)
	table.Render()
}
