package data

import (
	"billions_api/main/src/modules/wallet/models"
)

func AddWalletRepository(addWallet models.AddWallet) (models.Wallet, error) {

	wallet := models.Wallet{
		Name: addWallet.Name,
		Id:   "1",
	}

	return wallet, nil
}
