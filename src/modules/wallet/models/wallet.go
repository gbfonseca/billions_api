package models

type Wallet struct {
	Name string `json:"name" binding:"required"`
	Id   string `json:"id" binding:"required"`
}
