package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"time"
	"net/http"

	"github.com/valyala/fastjson"
)
type Bid struct {
	bidder string
	profile_id string
	amount int
	timestamp int
}

type Auction struct {
	uuid string 
	auctioneer string
	profile_id string
	start int
	end int
	item_name string
	item_lore string
	extra string
	category string
	tier string
	starting_bid int
	claimed bool
	highest_bid_amount int
	bids []Bid
}

func main() {

	start := time.Now()

	url := "https://api.hypixel.net/skyblock/auctions"

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "hypixel-auction-go")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//fmt.Printf("%s\n", body)

	v, err := fastjson.ParseBytes(body)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}

	a := v.GetArray("auctions")
	
	var auctions []Auction

	for i := 0; i < len(a); i++ {
		var auc Auction = Auction {
			uuid: string(a[i].GetStringBytes("uuid")),
			auctioneer: string(a[i].GetStringBytes("auctioneer")),
			profile_id: string(a[i].GetStringBytes("profile_id")),
			start: a[i].GetInt("start"),
			end: a[i].GetInt("end"),
			item_name: string(a[i].GetStringBytes("item_name")),
			item_lore: string(a[i].GetStringBytes("item_lore")),
			extra: string(a[i].GetStringBytes("extra")),
			category: string(a[i].GetStringBytes("category")),
			tier: string(a[i].GetStringBytes("tier")),
			starting_bid: a[i].GetInt("starting_bid"),
			claimed: a[i].GetBool("claimed"),
			highest_bid_amount: a[i].GetInt("highest_bid_amount"),
			bids: nil,
		}

		auctions = append(auctions, auc)
	}

	duration := time.Since(start)

	fmt.Println("Auction Item: ", auctions[0])

	fmt.Println("Execution Time: ", duration)
}