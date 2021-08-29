package twfxr

type Currency string

const (
	CurrencyUSD Currency = "USD" // 美金
	CurrencyHKD Currency = "HKD" // 港幣
	CurrencyGBP Currency = "GBP" // 英鎊
	CurrencyAUD Currency = "AUD" // 澳幣
	CurrencyCAD Currency = "CAD" // 加拿大幣
	CurrencySGD Currency = "SGD" // 新加坡幣
	CurrencyCHF Currency = "CHF" // 瑞士法郎
	CurrencyJPY Currency = "JPY" // 日圓
	CurrencyZAR Currency = "ZAR" // 南非幣
	CurrencySEK Currency = "SEK" // 瑞典幣
	CurrencyNZD Currency = "NZD" // 紐元
	CurrencyTHB Currency = "THB" // 泰幣
	CurrencyPHP Currency = "PHP" // 菲國比索
	CurrencyIDR Currency = "IDR" // 印尼幣
	CurrencyEUR Currency = "EUR" // 歐元
	CurrencyKRW Currency = "KRW" // 韓元
	CurrencyVND Currency = "VND" // 越南盾
	CurrencyMYR Currency = "MYR" // 越南盾
	CurrencyCNY Currency = "CNY" // 人民幣
)

type CurrencyExchangeRate struct {
	Currency string `json:"Currency"`
	// 本行買入
	BuyingCash           float64 `json:"Buying-Cash"` // 現金匯率
	BuyingSpot           float64 `json:"Buying-Spot"` // 即期匯率
	BuyingForward10Days  float64 `json:"Buying-Forward-10Days"`
	BuyingForward30Days  float64 `json:"Buying-Forward-30Days"`
	BuyingForward60Days  float64 `json:"Buying-Forward-60Days"`
	BuyingForward90Days  float64 `json:"Buying-Forward-90Days"`
	BuyingForward120Days float64 `json:"Buying-Forward-120Days"`
	BuyingForward150Days float64 `json:"Buying-Forward-150Days"`
	BuyingForward180Days float64 `json:"Buying-Forward-180Days"`
	// 本行賣出
	SellingCash           float64 `json:"Selling-Cash"` // 現金匯率
	SellingSpot           float64 `json:"Selling-Spot"` // 即期匯率
	SellingForward10Days  float64 `json:"Selling-Forward-10Days"`
	SellingForward30Days  float64 `json:"Selling-Forward-30Days"`
	SellingForward60Days  float64 `json:"Selling-Forward-60Days"`
	SellingForward90Days  float64 `json:"Selling-Forward-90Days"`
	SellingForward120Days float64 `json:"Selling-Forward-120Days"`
	SellingForward150Days float64 `json:"Selling-Forward-150Days"`
	SellingForward180Days float64 `json:"Selling-Forward-180Days"`
}
