package dtos

import (
	"time"
)

type Collection struct {
	Editors                 []string        `json:"editors"`
	PaymentTokens           []PaymentToken  `json:"payment_tokens"`
	PrimaryAssetContracts   []AssetContract `json:"primary_asset_contracts"`
	Traits                  interface{}
	Stats                   Stats
	BannerImageURL          string      `json:"banner_image_url"`
	ChatURL                 interface{} `json:"chat_url"`
	CreatedDate             time.Time   `json:"created_date"`
	DefaultToFiat           bool        `json:"default_to_fiat"`
	Description             string      `json:"description"`
	DevBuyerFeeBasisPoints  string      `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string      `json:"dev_seller_fee_basis_points"`
	DiscordURL              string      `json:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `json:"card_display_style"`
	} `json:"display_data"`
	ExternalURL                 string      `json:"external_url"`
	Featured                    bool        `json:"featured"`
	FeaturedImageURL            string      `json:"featured_image_url"`
	Hidden                      bool        `json:"hidden"`
	SafelistRequestStatus       string      `json:"safelist_request_status"`
	ImageURL                    string      `json:"image_url"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
	LargeImageURL               string      `json:"large_image_url"`
	MediumUsername              interface{} `json:"medium_username"`
	Name                        string      `json:"name"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address"`
	RequireEmail                bool        `json:"require_email"`
	ShortDescription            interface{} `json:"short_description"`
	Slug                        string      `json:"slug"`
	TelegramURL                 interface{} `json:"telegram_url"`
	TwitterUsername             string      `json:"twitter_username"`
	InstagramUsername           interface{} `json:"instagram_username"`
	WikiURL                     interface{} `json:"wiki_url"`
	IsNsfw                      bool        `json:"is_nsfw"`
	Fees                        struct {
		SellerFees  interface{} `json:"seller_fees"`
		OpenSeaFees interface{} `json:"opensea_fees"`
	} `json:"fees"`
	IsRarityEnabled bool `json:"is_rarity_enabled"`
}

type PaymentToken struct {
	Id       int     `json:"id"`
	Address  string  `json:"address"`
	ImageUrl string  `json:"image_url"`
	Name     string  `json:"Ether"`
	Decimals int     `json:"decimals"`
	EthPrice float32 `json:"eth_price"`
	UsdPrice float32 `json:"usd_price"`
}

type AssetContract struct {
	Address                     string `json:"address"`
	AssetContractType           string `json:"asset_contract_type"`
	CreatedDate                 string `json:"created_date"`
	Name                        string `json:"name"`
	NftVersion                  string `json:"nft_version"`
	OpenSeaVersion              string `json:"opensea_version"`
	Owner                       int    `json:"owner"`
	SchemaName                  string `json:"schema_name"`
	Symbol                      string `json:"symbol"`
	TotalSupply                 string `json:"total_supply"`
	Description                 string `json:"description"`
	ExternalLink                string `json:"external_link"`
	ImageUrl                    string `json:"image_url"`
	DefaultToFiat               bool   `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int    `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int    `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenSeaBuyerFeeBasisPoints  int    `json:"opensea_buyer_fee_basis_points"`
	OpenSeaSellerFeeBasisPoints int    `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int    `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int    `json:"seller_fee_basis_points"`
	PayoutAddress               string `json:"payout_address"`
}

type Stats struct {
	OneHourVolume         float64 `json:"one_hour_volume"`
	OneHourChange         float64 `json:"one_hour_change"`
	OneHourSales          float64 `json:"one_hour_sales"`
	OneHourSalesChange    float64 `json:"one_hour_sales_change"`
	OneHourAveragePrice   float64 `json:"one_hour_average_price"`
	OneHourDifference     float64 `json:"one_hour_difference"`
	SixHourVolume         float64 `json:"six_hour_volume"`
	SixHourChange         float64 `json:"six_hour_change"`
	SixHourSales          float64 `json:"six_hour_sales"`
	SixHourSalesChange    float64 `json:"six_hour_sales_change"`
	SixHourAveragePrice   float64 `json:"six_hour_average_price"`
	SixHourDifference     float64 `json:"six_hour_difference"`
	OneDayVolume          float64 `json:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change"`
	OneDaySales           float64 `json:"one_day_sales"`
	OneDaySalesChange     float64 `json:"one_day_sales_change"`
	OneDayAveragePrice    float64 `json:"one_day_average_price"`
	OneDayDifference      float64 `json:"one_day_difference"`
	SevenDayVolume        float64 `json:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change"`
	SevenDaySales         float64 `json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
	SevenDayDifference    float64 `json:"seven_day_difference"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change"`
	ThirtyDaySales        float64 `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
	ThirtyDayDifference   float64 `json:"thirty_day_difference"`
	TotalVolume           float64 `json:"total_volume"`
	TotalSales            float64 `json:"total_sales"`
	TotalSupply           float64 `json:"total_supply"`
	Count                 float64 `json:"count"`
	NumOwners             int     `json:"num_owners"`
	AveragePrice          float64 `json:"average_price"`
	NumReports            int     `json:"num_reports"`
	MarketCap             float64 `json:"market_cap"`
	FloorPrice            float64 `json:"floor_price"`
}
