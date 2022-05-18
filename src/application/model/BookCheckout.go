package model

type BookCheckout struct {
	BlockID      string `json:"block_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}
