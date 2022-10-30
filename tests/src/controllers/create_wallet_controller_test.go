package src

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"billions_api/main/src/modules/wallet/controllers"
	"billions_api/main/src/modules/wallet/models"
	fixtures_test "billions_api/main/tests/fixtures"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnBadRequestOnInvalidBody(t *testing.T) {

	w := httptest.NewRecorder()
	addWallet := models.AddWallet{}
	jsonValue, _ := json.Marshal(addWallet)

	ctx := fixtures_test.GetTestGinContext(w)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonValue))

	controllers.CreateWalletController(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestShouldReturnAnWalletOnSuccess(t *testing.T) {

	w := httptest.NewRecorder()
	addWallet := models.AddWallet{
		Name: "My new wallet",
	}
	jsonValue, _ := json.Marshal(addWallet)

	ctx := fixtures_test.GetTestGinContext(w)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonValue))

	controllers.CreateWalletController(ctx)

	var wallet models.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, wallet.Id)
}
