package bitflyer

import (
	"buy-btc/utils"
	"encoding/json"
)

const baseURL = "https://api.bitflyer.com/v1/ticker/?product_code=BTC_JPY"
const productCodeKey = "product_code"

func GetTicker(code ProductCode) (*Ticker, error) {
	url := baseURL + "/v1/ticker"
	res, err := utils.DoHttpRequest("GET", url, nil, map[string]string{productCodeKey: code.String()}, nil)
	if err != nil {
		return nil, err
	}

	var ticker Ticker
	err = json.Unmarshal(res, &ticker)
	if err != nil {
		return nil, err
	}
	return &ticker, nil
}

type Ticker struct {
	ProductCode       string  `json:"product_code"`
	State             string  `json:"state"`
	Timestamp         string  `json:"timestamp"`
	Tick_id           int     `json:"tick_id"`
	Best_bid          float64 `json:"best_bid"`
	Best_ask          float64 `json:"best_ask"`
	Best_bid_size     float64 `json:"best_bid_size"`
	Best_ask_size     float64 `json:"best_ask_size"`
	Total_bid_depth   float64 `json:"total_bid_depth"`
	Total_ask_depth   float64 `json:"total_ask_depth"`
	Market_bid_size   float64 `json:"market_bid_size"`
	Market_ask_size   float64 `json:"market_ask_size"`
	Ltp               float64 `json:"ltp"`
	Volume            float64 `json:"volume"`
	Volume_by_product float64 `json:"volume_by_product"`
}
