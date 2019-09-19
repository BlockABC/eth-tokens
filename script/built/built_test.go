package built

import (
	"fmt"
	"testing"
)

func TestReadTokenIcon(t *testing.T) {
	dir := "/Users/plainchant/go/src/github.com/eager7/eth_tokens/tokens/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
	token, err := ReadTokenInfo(dir)
	if err != nil {
		fmt.Println("read token info err:", err)
		t.Fatal(err)
	}
	if err := ReadTokenIcon(dir, token); err != nil {
		t.Fatal(err)
	}
}
