package coin_gecko

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eager7/elog"
	"github.com/BlockABC/eth-tokens/script/built"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"net/http"
	"net/url"
)

var log = elog.NewLogger("gecko", elog.DebugLevel)

const urlGecko = "https://api.coingecko.com/api/v3/coins/ethereum/contract/%s"

type CoinGecko struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Image  struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
}

func ReplaceTokenLogoFromCoinGecko(token *built.TokenInfo) {
	var err error
	for i := 0; i < 3; i++ {
		if err := RequestTokenInfoFromCoinGecko(token); err == nil {
			return
		}
	}
	log.Error("RequestTokenInfoFromCoinGecko err:", err, token.Address)
}

func RequestTokenInfoFromCoinGecko(token *built.TokenInfo) error {
	u := fmt.Sprintf(urlGecko, common.HexToAddress(token.Address).Hex())
	log.Notice("request url:", u)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return errors.New("http request err:" + err.Error())
	}
	q := url.Values{}
	//q.Add("limit", "5000")
	req.Header.Set("accept", "application/json")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("client do err:" + err.Error())
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println("body close err:", err)
		}
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("read all err:" + err.Error())
	}
	var coin CoinGecko
	if err := json.Unmarshal(body, &coin); err != nil {
		return errors.New("json unmarshal err:" + err.Error())
	}
	token.Logo.Src = coin.Image.Large
	if token.Name == "" {
		token.Name = coin.Name
	}
	if token.Symbol == "" {
		token.Symbol = coin.Symbol
	}
	return nil
}
