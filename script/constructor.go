package main

import (
	"encoding/json"
	"fmt"
	"github.com/BlockABC/wallet_eth_webserver/request"
	"reflect"
	"strings"
)

const urlTokenList = "https://raw.githubusercontent.com/MyEtherWallet/ethereum-lists/master/dist/tokens/eth/tokens-eth.json"

type Logo struct {
	Src      string      `json:"src"`
	Width    interface{} `json:"width"` //这两个字段在返回中有的是字符，有的是数字，需要进行二次解析
	Height   interface{} `json:"height"`
	IpfsHash string      `json:"ipfs_hash"`
}
type Support struct {
	Email string `json:"email"`
	Url   string `json:"url"`
}

type Social struct {
	Blog      string `json:"blog"`
	Chat      string `json:"chat"`
	Facebook  string `json:"facebook"`
	Forum     string `json:"forum"`
	Github    string `json:"github"`
	Gitter    string `json:"gitter"`
	Instagram string `json:"instagram"`
	Linkedin  string `json:"linkedin"`
	Reddit    string `json:"reddit"`
	Slack     string `json:"slack"`
	Telegram  string `json:"telegram"`
	Twitter   string `json:"twitter"`
	Youtube   string `json:"youtube"`
}
type TokenInfo struct {
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Address    string  `json:"address"`
	EnsAddress string  `json:"ens_address"`
	Decimals   int     `json:"decimals"`
	Website    string  `json:"website"`
	Logo       Logo    `json:"logo"`
	Support    Support `json:"support"`
	Social     Social  `json:"social"`
}

func TokenListFromGit(url string) (string, error) {
	var tokenLists []TokenInfo
	body, err := request.GetResponseBytes(request.NewRequester(url))
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &tokenLists); err != nil {
		return "", err
	}
	for _, token := range tokenLists {
		switch reflect.TypeOf(token.Logo.Width).Kind() {
		case reflect.Float64:
			token.Logo.Width = fmt.Sprintf("%v", token.Logo.Width.(float64))
			token.Logo.Height = fmt.Sprintf("%v", token.Logo.Height.(float64))
		default:
			continue
		}
	}
	ret, err := json.Marshal(tokenLists)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func main() {
	fmt.Println(strings.ToLower("0xB8c77482e45F1F44dE1745F52C74426C631bDD52"))
}
