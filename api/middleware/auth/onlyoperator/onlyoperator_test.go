package onlyoperator

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/ethwallet"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func Test_OnlyOperator(t *testing.T) {

	t.Run("should fail if wallet is not of operator", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		walletAddress := testingcommon.GenerateWallet().WalletAddress
		c.Set("walletAddress", walletAddress)
		OnlyOperator(c)
		assert.Equal(t, http.StatusForbidden, rr.Result().StatusCode)
	})

	t.Run("should pass if wallet is of operator", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		_, pubKy, _, err := ethwallet.HdWallet(envutil.MustGetEnv("MNEMONIC"))
		if err != nil {
			t.Fatal(err)
		}
		walletAddress := crypto.PubkeyToAddress(*pubKy).String()
		c.Set("walletAddress", walletAddress)
		OnlyOperator(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})
}
