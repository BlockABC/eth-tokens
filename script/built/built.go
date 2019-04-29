package built

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Contract string `json:"contract"`
	Decimals int    `json:"decimals"`
	Logo     string `json:"logo"`
	Desc     struct {
		En string `json:"en"`
		Zh string `json:"zh"`
	} `json:"desc"`

	WebSite    string `json:"website"`
	WhitePaper string `json:"whitepaper"`
	Invalid    bool   `json:"invalid"`
	Links      struct {
		Twitter  string `json:"twitter"`
		Telegram string `json:"telegram"`
	} `json:"links"`
}

func BuildDist(dir string, tokens []*Token) error {
	file, err := os.OpenFile(dir+"/tokens.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.MarshalIndent(tokens, "", "    ")
	if err != nil {
		return err
	}
	_, err = file.WriteString(string(data))
	if err != nil {
		return err
	}
	return nil
}

func BuildReadme(fil string, tokens []*Token) error {
	var tokensInfo []string
	for _, token := range tokens {
		tokensInfo = append(tokensInfo, fmt.Sprintf(`|  <img src="%s" width=30 />  | [%s](https://github.com/eager7/eth_tokens/blob/master/tokens/%s/%s.json) | [%s](https://etherscan.io/address/%s) |`,
			token.Logo, token.Symbol, strings.ToLower(token.Contract), FormatSymbol(token.Symbol), token.Contract, token.Contract))
	}
	file, err := os.OpenFile(fil, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(`## Token List
<!-- token_list_start -->
|   Logo    | Symbol      | Account Name |
| ----------- |:------------:|:------------:|` + "\n")
	for _, s := range tokensInfo {
		_, err = file.WriteString(s + "\n")
	}
	return nil
}

func CollectTokens(dir string) (tokens []*Token, err error) {
	dirList, err := TokensDirList(dir)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirList {
		token, err := ReadTokenInfo(dir)
		if err != nil {
			fmt.Println("read token info err:", err)
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func ReadTokenInfo(dir string) (*Token, error) {
	names, err := filepath.Glob(dir + "/*.json")
	if err != nil {
		return nil, err
	}
	if len(names) != 1 {
		return nil, errors.New(fmt.Sprintf("file format wrong:%s", dir))
	}
	file, err := os.Open(names[0])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open %s err:%s", dir, err.Error()))
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("read %s err:%s", file.Name(), err.Error()))
	}
	var token Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, errors.New("json un marshal err:" + err.Error())
	}
	return &token, nil
}

func TokensDirList(dir string) ([]string, error) {
	var dirList []string
	err := filepath.Walk(dir,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dirList = append(dirList, path)
				return nil
			}
			return nil
		})
	return dirList[1:], err
}
