package coin_gecko

import (
	"github.com/eager7/eth_tokens/script/built"
	"testing"
)

func TestRequestLogoFromCoinGecko(t *testing.T) {
	token := built.Token{
		Contract: "0x0000000000085d4780B73119b644AE5ecd22b376",
	}
	if err := RequestTokenInfoFromCoinGecko(&token); err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
