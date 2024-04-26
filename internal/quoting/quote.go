package quoting

import "time"

type Quote struct {
	ID                           string  `json:"id"`
	Symbol                       string  `json:"symbol"`
	Name                         string  `json:"name"`
	AssetPlatformID              any     `json:"asset_platform_id"`
	SentimentVotesUpPercentage   float64 `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64 `json:"sentiment_votes_down_percentage"`
	WatchlistPortfolioUsers      int     `json:"watchlist_portfolio_users"`
	MarketCapRank                int     `json:"market_cap_rank"`
	MarketData                   struct {
		CurrentPrice struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"current_price"`
		MarketCap struct {
			Cad int64 `json:"cad"`
			Usd int64 `json:"usd"`
		} `json:"market_cap"`
		MarketCapRank                int     `json:"market_cap_rank"`
		PriceChange24H               float64 `json:"price_change_24h"`
		PriceChangePercentage24H     float64 `json:"price_change_percentage_24h"`
		PriceChangePercentage7D      float64 `json:"price_change_percentage_7d"`
		PriceChangePercentage14D     float64 `json:"price_change_percentage_14d"`
		PriceChangePercentage30D     float64 `json:"price_change_percentage_30d"`
		PriceChangePercentage60D     float64 `json:"price_change_percentage_60d"`
		PriceChangePercentage200D    float64 `json:"price_change_percentage_200d"`
		PriceChangePercentage1Y      float64 `json:"price_change_percentage_1y"`
		MarketCapChange24H           float64 `json:"market_cap_change_24h"`
		MarketCapChangePercentage24H float64 `json:"market_cap_change_percentage_24h"`
		PriceChange24HInCurrency     struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_24h_in_currency"`
		PriceChangePercentage1HInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_percentage_1h_in_currency"`
		PriceChangePercentage24HInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_percentage_24h_in_currency"`
		PriceChangePercentage7DInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_percentage_7d_in_currency"`
		PriceChangePercentage14DInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_percentage_14d_in_currency"`
		PriceChangePercentage30DInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_percentage_30d_in_currency"`
		PriceChangePercentage60DInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_percentage_60d_in_currency"`
		PriceChangePercentage200DInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_percentage_200d_in_currency"`
		PriceChangePercentage1YInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"price_change_percentage_1y_in_currency"`
		MarketCapChange24HInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"market_cap_change_24h_in_currency"`
		MarketCapChangePercentage24HInCurrency struct {
			Cad float64 `json:"cad"`
			Usd float64 `json:"usd"`
		} `json:"market_cap_change_percentage_24h_in_currency"`
		TotalSupply       float64   `json:"total_supply"`
		MaxSupply         any       `json:"max_supply"`
		CirculatingSupply float64   `json:"circulating_supply"`
		LastUpdated       time.Time `json:"last_updated"`
	} `json:"market_data"`
}
