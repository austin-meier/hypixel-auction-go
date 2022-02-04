package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/valyala/fastjson"
)

type Bid struct {
	bidder    string
	profileId string
	amount    int
	timestamp int
}

type Auction struct {
	uuid             string
	auctioneer       string
	profileId        string
	start            int
	end              int
	itemName         string
	itemLore         string
	extra            string
	category         string
	tier             string
	startingBid      int
	claimed          bool
	highestBidAmount int
	bids             []Bid
}

func request(page int, client http.Client) []*fastjson.Value {
	url := "https://api.hypixel.net/skyblock/auctions?page=" + strconv.Itoa(page)
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
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
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

	return v.GetArray("auctions")

}

func newClient() http.Client {
	tr := &http.Transport{
		MaxIdleConnsPerHost: 100,
	}

	client := http.Client{
		Timeout:   time.Second * 100, // Timeout after 2 seconds
		Transport: tr,                // Keeps connections open
	}

	return client
}

func main() {

	start := time.Now()

	client := newClient()

	a := request(0, client)

	var auctions []Auction

	for i := 0; i < len(a); i++ {
		var auc Auction = Auction{
			uuid:             string(a[i].GetStringBytes("uuid")),
			auctioneer:       string(a[i].GetStringBytes("auctioneer")),
			profileId:        string(a[i].GetStringBytes("profile_id")),
			start:            a[i].GetInt("start"),
			end:              a[i].GetInt("end"),
			itemName:         string(a[i].GetStringBytes("item_name")),
			itemLore:         string(a[i].GetStringBytes("item_lore")),
			extra:            string(a[i].GetStringBytes("extra")),
			category:         string(a[i].GetStringBytes("category")),
			tier:             string(a[i].GetStringBytes("tier")),
			startingBid:      a[i].GetInt("starting_bid"),
			claimed:          a[i].GetBool("claimed"),
			highestBidAmount: a[i].GetInt("highest_bid_amount"),
			bids:             nil,
		}

		auctions = append(auctions, auc)
	}

	fmt.Println("Auction Item: ", auctions[0])

	duration := time.Since(start)

	fmt.Println("Execution Time: ", duration)

}
