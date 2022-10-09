package src

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"billions_api/main/src/controllers"
	wallet "billions_api/main/src/domain/usecases/wallet"
	fixtures_test "billions_api/main/tests/fixtures"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnBadRequestOnInvalidBody(t *testing.T) {

	w := httptest.NewRecorder()
	addWallet := wallet.AddWallet{}
	jsonValue, _ := json.Marshal(addWallet)

	ctx := fixtures_test.GetTestGinContext(w)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonValue))

	controllers.CreateWalletController(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}