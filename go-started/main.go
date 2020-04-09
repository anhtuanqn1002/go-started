package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strings"

	"github.com/anhtuanqn1002/redis"
)

// Coin struct
type Coin struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func getBTCprice(symbol string) string {
	url := "https://api.binance.com/api/v3/ticker/price?symbol=" + strings.ToUpper(symbol) + "BTC"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var coin Coin
	json.Unmarshal(body, &coin)
	return coin.Price
}

func main() {
	client := redis.RClient()

	// check connection status
	err := redis.Ping(client)
	if err != nil {
		fmt.Println(err)
	}

	// Using the SET command to set Key-value pair
	// err = redis.Set(client)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	var name = "name"
	err = redis.Set(client, name, "Lin")
	if err != nil {
		fmt.Println(err)
	}

	value, err := redis.Get(client, name)
	if err == nil {
		fmt.Println(value)
	}

	// Using the GET command to get values from keys
	// err = redis.Get(client)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
