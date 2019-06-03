package main

import (
	"flag"
	"fmt"
	"github.com/eager7/eth_tokens/script/built"
	"github.com/eager7/eth_tokens/script/coin_gecko"
	"github.com/eager7/eth_tokens/script/ether_scan"
)

func main() {
	var g = flag.Bool("g", false, "")
	var e = flag.Bool("e", false, "")
	var logo = flag.Bool("logo", false, "")
	flag.Parse()
	if *g {
		tokenList, err := built.TokenListFromGit(built.URLTokenList)
		if err != nil {
			panic(err)
		}
		fmt.Println("success get token list:", len(tokenList))
		if err := built.InitializeTokens(`../../tokens`, tokenList); err != nil {
			panic(err)
		}
	}
	if *e {
		spider, err := ether_scan.Initialize("http://47.52.157.31:8585")
		if err != nil {
			panic(err)
		}
		tokenList, err := spider.BuiltTokensFromEtherScan()
		if err != nil {
			panic(err)
		}
		fmt.Println("success get token list:", len(tokenList))
		if err := built.InitializeTokens(`../../tokens`, tokenList); err != nil {
			panic(err)
		}
	}

	tokens, err := built.CollectTokens(`../../tokens`)
	if err != nil {
		panic(err)
	}

	if *logo {
		for _, token := range tokens {
			coin_gecko.ReplaceTokenLogoFromCoinGecko(token)
		}
	}
	eth := built.Token{
		Name:     "Ethereum",
		Symbol:   "ETH",
		Contract: "0x0000000000000000000000000000000000000000",
		Decimals: 18,
		Logo:     "https://www.cryptocompare.com/media/20646/eth_logo.png?width=200",
		Invalid:  true,
	}
	tokens = append([]*built.Token{&eth}, tokens...)
	if err := built.BuildDist(`../../dist`, tokens); err != nil {
		panic(err)
	}
	if err := built.BuildReadme(`../../tokens.md`, tokens); err != nil {
		panic(err)
	}
}
