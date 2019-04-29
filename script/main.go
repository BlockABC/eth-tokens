package main

import "github.com/eager7/eth_tokens/script/built"

func main() {
	tokenList, err := built.TokenListFromGit(built.URLTokenList)
	if err != nil {
		panic(err)
	}
	if err := built.InitializeTokens(`../tokens`, tokenList); err != nil {
		panic(err)
	}

	tokens, err := built.CollectTokens(`../tokens`)
	if err != nil {
		panic(err)
	}
	if err := built.BuildDist(`../dist`, tokens); err != nil {
		panic(err)
	}
	if err := built.BuildReadme(`../tokens.md`, tokens); err != nil {
		panic(err)
	}
}
