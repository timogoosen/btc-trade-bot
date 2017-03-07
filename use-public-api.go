package main

import (
	"fmt"
	"github.com/bitx/bitx-go"
	"log"
)




func main() {

	// curl https://api.mybitx.com/api/1/trades?pair=XBTZAR

	// Pass an empty string for the api_key_id if you will only access the public
	// API.
	//func NewClient(api_key_id, api_key_secret string) *Client

	pair := "XBTZAR"
	c := bitx.NewClient("", "test")

	trade, err := c.Trades(pair)

	if err != nil {
		log.Fatal(err)

	}

	// Loop over Trade array
	for i := 0; i < len(trade); i++ {
		fmt.Println()

		fmt.Println("Timestamp:",trade[i].Timestamp)
		fmt.Println("Price:",trade[i].Price)
		fmt.Println("Volume:",trade[i].Volume)
		total_transaction_cost := float32(trade[i].Price * trade[i].Volume)
		fmt.Println("Total Transaction Cost:", total_transaction_cost)

		// Log all of this to a temporary SQLITE Table.

		// We when use this table to find the largest transaction.

		// The largest transaction gets INSERTED To a table (not temporary), then we drop the temporary table.

		

	}

}
