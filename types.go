package dexscreenerapi

type TokenProfile struct {
	URL          string `json:"url"`
	ChainID      string `json:"chainId"`
	TokenAddress string `json:"tokenAddress"`
	Icon         string `json:"icon"`
	Header       string `json:"header"`
	Description  string `json:"description"`
	Links        []Link `json:"links"`
}

type BoostedTokenProfile struct {
	URL          string  `json:"url"`
	ChainID      string  `json:"chainId"`
	TokenAddress string  `json:"tokenAddress"`
	Amount       float64 `json:"amount"`
	TotalAmount  float64 `json:"totalAmount"`
	Icon         string  `json:"icon"`
	Header       string  `json:"header"`
	Description  string  `json:"description"`
	Links        []Link  `json:"links"`
}

type Link struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

type PairsResponse struct {
	SchemaVersion string `json:"schemaVersion"`
	Pairs         []Pair `json:"pairs"`
}

type Pair struct {
	ChainID       string    `json:"chainId"`
	DexID         string    `json:"dexId"`
	URL           string    `json:"url"`
	PairAddress   string    `json:"pairAddress"`
	Labels        []string  `json:"labels"`
	BaseToken     Token     `json:"baseToken"`
	QuoteToken    Token     `json:"quoteToken"`
	PriceNative   string    `json:"priceNative"`
	PriceUsd      string    `json:"priceUsd"`
	Liquidity     Liquidity `json:"liquidity"`
	FDV           float64   `json:"fdv"`
	MarketCap     float64   `json:"marketCap"`
	PairCreatedAt int64     `json:"pairCreatedAt"`
	Info          Info      `json:"info"`
	Boosts        Boosts    `json:"boosts"`
}

type Token struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

type Liquidity struct {
	USD   float64 `json:"usd"`
	Base  float64 `json:"base"`
	Quote float64 `json:"quote"`
}

type Info struct {
	ImageURL string    `json:"imageUrl"`
	Websites []Website `json:"websites"`
	Socials  []Social  `json:"socials"`
}

type Website struct {
	URL string `json:"url"`
}

type Social struct {
	Platform string `json:"platform"`
	Handle   string `json:"handle"`
}

type Boosts struct {
	Active int `json:"active"`
}

type OrderPaidResponse struct {
	Type             string `json:"type"`
	Status           string `json:"status"`
	PaymentTimestamp int    `json:"paymentTimestamp"`
}
