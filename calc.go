package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "os"
  "encoding/json"
  "math"
  // "reflect"
)

type CryptoPrice struct {
    Status Status `json:"status"`
    Data   Data   `json:"data"`
}
type Status struct {
    Timestamp    string  `json:"timestamp"`
    ErrorCode    int         `json:"error_code"`
    ErrorMessage interface{} `json:"error_message"`
    Elapsed      int         `json:"elapsed"`
    CreditCount  int         `json:"credit_count"`
    Notice       interface{} `json:"notice"`
}
type USD struct {
    Price       float64   `json:"price"`
    LastUpdated string `json:"last_updated"`
}
type Quote struct {
    USD USD `json:"USD"`
}
type Data struct {
    ID          int       `json:"id"`
    Symbol      string    `json:"symbol"`
    Name        string    `json:"name"`
    Amount      int       `json:"amount"`
    LastUpdated string `json:"last_updated"`
    Quote       Quote     `json:"quote"`
}

func price(symbol string) float64 {
  
  symb := symbol

  client := &http.Client{}
  req, err := http.NewRequest("GET","https://pro-api.coinmarketcap.com/v1/tools/price-conversion", nil)
  if err != nil {
    log.Print(err)
    os.Exit(1)
  }

  q := url.Values{}
  
  q.Add("amount", "1")
  q.Add("symbol", symb)
  q.Add("convert", "USD")

  req.Header.Set("Accepts", "application/json")
  req.Header.Add("X-CMC_PRO_API_KEY", "84b9e7f8-3a02-4a21-841f-3787098f99c5")
  req.URL.RawQuery = q.Encode()


  resp, err := client.Do(req);
  if err != nil {
    fmt.Println("Error sending request to server")
    os.Exit(1)
  }
  // fmt.Println(resp.Status);
  respBody, _ := ioutil.ReadAll(resp.Body)
  response := string(respBody);
  // fmt.Println(response);

  var prices CryptoPrice 
  json.Unmarshal([]byte(response), &prices)
  value := prices.Data.Quote.USD.Price
  // fmt.Println(btcPrice)
  return math.Round(value*100)/100
}
// func btcPrice() {
//     price("BTC")
//     //fmt.Println(price("BTC"))

// }
// func ethPrice() {
//     // price("BTC")
//     fmt.Println(price("ETH"))

// }
func Mul(param1 float64, param2 float64) float64 {
    return math.Round((param1 * param2)*1000)/1000
}