package main

import (
	"database/sql"
	"fmt"
	"github.com/bitx/bitx-go"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
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

	// Log all of this to a temporary SQLITE Table.
	// We when use this table to find the largest transaction.
	// The largest transaction gets INSERTED To a table (not temporary), then we drop the temporary table.
	// Remove db file if it was craeted before
	os.Remove("tempdb.db")
	db, err := sql.Open("sqlite3", "./tempdb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// create table temptable (id integer not null primary key, timestamp text, price real, volume real, total_transaction_cost real);
	// See if id gets generated automatically if we don't specify it explicitly
	sqlStmt := `
 create table temptable (id integer not null primary key,timestamp text, price real, volume real, total_transaction_cost real);
 delete from temptable;
 `
	// Not sure what the delete statement is for I assume clear the table so to make sure it is empty?
	// Got this example from : https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	//  stmt, err := tx.Prepare("insert into temptable(id, name) values(?, ?)")

	// Loop over Trade array
	for i := 0; i < len(trade); i++ {
		fmt.Println()

		fmt.Println("Timestamp:", trade[i].Timestamp)
		fmt.Println("Price:", trade[i].Price)
		fmt.Println("Volume:", trade[i].Volume)
		total_transaction_cost := float32(trade[i].Price * trade[i].Volume)
		fmt.Println("Total Transaction Cost:", total_transaction_cost)

		stmt, err := tx.Prepare("insert into temptable(id, timestamp, price, volume, total_transaction_cost) values(?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(i, trade[i].Timestamp, trade[i].Price, trade[i].Volume, total_transaction_cost)
		if err != nil {
			log.Fatal(err)
		}


	}
		tx.Commit()
		fmt.Println("Commit outside of for loop. Hope this works!")

}
