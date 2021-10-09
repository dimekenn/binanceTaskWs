package configs

type Config struct {
	//server port
	Port string `json:"port"`
	//binance stream url
	WsUrl string `json:"ws_url"`
	//orders count in each book
	OrdersInBook int8 `json:"orders_in_book"`
	//time range between messages
	TimeRange string `json:"time_range"`
}

func NewConfig() *Config {
	return &Config{}
}
