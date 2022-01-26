package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"time"
	"net/http"

	"github.com/valyala/fastjson"
)

func main() {

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

	// v, err := fastjson.ParseBytes(body)
	// if err != nil {
	// 	log.Fatalf("cannot parse json: %s", err)
	// }

	v:= fastjson.GetString(body, "auctions", "1", "item_name")

	fmt.Printf("%s\n", v)
}