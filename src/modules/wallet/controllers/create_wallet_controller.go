package controllers

import (
	"net/http"

	"billions_api/main/src/modules/wallet/models"

	"github.com/gin-gonic/gin"
)

func CreateWalletController(ctx *gin.Context) {

	var requestBody models.AddWallet

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"name": "123",
		"id":   "1",
	})

}
