package coin_gecko

import (
	"github.com/BlockABC/eth-tokens/script/built"
	"testing"
)

func TestRequestLogoFromCoinGecko(t *testing.T) {
	token := built.TokenInfo{
		Address: "0x0000000000085d4780B73119b644AE5ecd22b376",
	}
	if err := RequestTokenInfoFromCoinGecko(&token); err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
