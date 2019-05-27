package erc20

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestReadTokenInfo(t *testing.T) {
	client, err := ethclient.Dial("http://47.52.157.31:8585")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ReadTokenInfo(`0x6aedbf8dff31437220df351950ba2a3362168d1b`, client, "http://47.52.157.31:8585"))
}
