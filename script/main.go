package main

import (
	"flag"
	"fmt"
	"github.com/eager7/eth_tokens/script/built"
)

func main() {
	var g = flag.Bool("g", false, "")
	flag.Parse()
	if *g {
		tokenList, err := built.TokenListFromGit(built.URLTokenList)
		if err != nil {
			panic(err)
		}
		fmt.Println("success get token list:", len(tokenList))
		if err := built.InitializeTokens(`../tokens`, tokenList); err != nil {
			panic(err)
		}
	}

	tokens, err := built.CollectTokens(`../tokens`)
	if err != nil {
		panic(err)
	}
	eth := built.Token{
		Name:     "Ethereum",
		Symbol:   "ETH",
		Contract: "0x0000000000000000000000000000000000000000",
		Decimals: 18,
		Logo:     "https://raw.githubusercontent.com/eager7/eth_tokens/master/tokens/0x0000000000000000000000000000000000000000/token.png",
		Invalid:  true,
	}
	tokens = append([]*built.Token{&eth}, tokens...)
	if err := built.BuildDist(`../dist`, tokens); err != nil {
		panic(err)
	}
	if err := built.BuildReadme(`../tokens.md`, tokens); err != nil {
		panic(err)
	}
}
