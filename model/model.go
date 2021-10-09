package model

type StrArr []string

type Msg struct {
	LastUpdateId int `json:"lastUpdateId"`
	Bids []StrArr `json:"bids"`
	Asks []StrArr `json:"asks"`
}
