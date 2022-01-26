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
	
	text, err := fastjson.Parser.Parse(`{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}

	fmt.Printf("%s\n", text)
	
}