package bitflyer

type ProductCode int

const (
	Btcjpy ProductCode = iota
	Ethjpy
	Exbtcjpy
	Ethbtc
	bchbtc
)

func (code ProductCode) string() string {
	switch code {
	case Btcjpy:
		return "BTC_JPY"
	case Ethjpy:
		return "ETH_JPY"
	case Exbtcjpy:
		return "FX_BTC_JPY"
	case Ethbtc:
		return "ETH_JPY"
	case bchbtc:
		return "BCH_JPY"
	default:
		return "BTC_JPY"
	}
}
