package ether_scan

import (
	"fmt"
	"github.com/eager7/eth_tokens/script/built"
	"regexp"
	"testing"
)

func TestRequestErc20ListByPage(t *testing.T) {
	tokens, err := RequestTokenListByPage(urlEtherScan + "10")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("tokens:", len(tokens))
}

func TestRequestNftListByPage(t *testing.T) {
	tokens, err := RequestTokenListByPage(urlNftEtherScan + "1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("tokens:", len(tokens))
}

func TestBuilt(t *testing.T) {
	spider, err := Initialize("http://47.52.157.31:8585")
	if err != nil {
		panic(err)
	}
	tokenList, err := spider.BuiltTokensFromEtherScan()
	if err != nil {
		panic(err)
	}
	fmt.Println("success get token list:", len(tokenList))
	for i, token := range tokenList {
		fmt.Println(i, "-------------")
		fmt.Println(token.Address)
		fmt.Println(token.Logo.Src)
	}
}

func TestReg(t *testing.T) {
	ret := `Bottos is a platform for the value exchange of AI and affiliated industries based on data feeding, with the ultimate goal of building a distributed AI new ecosystem.</p></div> <img class="u-xs-avatar mr-2" src="/token/images/bottos_28_3.png"/><div class="media-body"><h3 class="h6 mb-0"><a class="text-primary" href="/token/0x36905fc93280f52362a1cbab151f25dc46742fb5">BTOCoin (BTO)</a></h3><p class="d-none d-md-block font-size-1 mb-0">Bottos is more than an AI public chain with both data
and model marketplaces,
Bottos is a platform for the value exchange of AI and affiliated industries based on data feeding, with the ultimate goal of building a distributed AI new ecosystem.</p></div>`

	reIcon := regexp.MustCompile(`.*?src="(?P<icon>[^"]*)(?s:".*)`)
	reContract := regexp.MustCompile(`.*?href="/token/(?P<contract>[^"]*)(?s:".*)`)
	icon := reIcon.ReplaceAllString(ret, "$icon")
	contract := reContract.ReplaceAllString(ret, "$contract")
	fmt.Println("--------------------------------------")
	fmt.Println(icon)
	fmt.Println("--------------------------------------")
	fmt.Println(contract)
}

func TestRequestTokenLogo(t *testing.T) {
	token := built.TokenInfo{
		Address: `0x501262281b2ba043e2fbf14904980689cddb0c78`,
	}
	if err := RequestTokenLogo(&token); err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestBuiltNftFromEtherScan(t *testing.T) {
	spider, err := Initialize("http://47.106.254.3:8585")
	if err != nil {
		panic(err)
	}
	tokenList, err := spider.BuiltNftFromEtherScan()
	if err != nil {
		panic(err)
	}
	fmt.Println("success get token list:", len(tokenList))
	for i, token := range tokenList {
		fmt.Println(i, "-------------")
		fmt.Println(token.Address)
		fmt.Println(token.Logo.Src)
	}
}