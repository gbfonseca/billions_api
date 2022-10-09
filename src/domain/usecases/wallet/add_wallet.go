package usecases

type AddWallet struct {
	Name string `json:"name" binding:"required"`
}
