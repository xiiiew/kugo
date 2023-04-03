package test

import (
	"github.com/shopspring/decimal"
	"github.com/xiiiew/kugo"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var instance *kugo.Kucoin

const (
	spotEndpoint = "https://openapi-sandbox.kucoin.com"
	accessKey    = "accessKey"
	secretKey    = "secretKey"
	passphrase   = "12345678"

	futureEndpoint   = "https://api-sandbox-futures.kucoin.com"
	futureAccessKey  = "accessKey"
	futureSecretKey  = "secretKey"
	futurePassphrase = "12345678"
)

func TestMain(m *testing.M) {
	uProxy, _ := url.Parse("http://127.0.0.1:7890")
	i, err := kugo.NewKucoin(
		kugo.SetSpotEndpoint(spotEndpoint),
		kugo.SetFutureEndpoint(futureEndpoint),
		kugo.SetApiKey(accessKey, secretKey, passphrase),
		kugo.SetDebug(true),
		kugo.SetClient(&http.Client{Transport: &http.Transport{
			Proxy:             http.ProxyURL(uProxy),
			DisableKeepAlives: true},
			Timeout: 10 * time.Second}),
	)
	if err != nil {
		log.Fatalln(err)
	}

	instance = i
	m.Run()
}

func TestSymbols(t *testing.T) {
	symbols, err := instance.Symbols("USDS")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", symbols)
}

func TestAccounts(t *testing.T) {
	accounts, err := instance.Accounts("BTC", "trade")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", accounts)
}

func TestSpotOrder(t *testing.T) {
	req := &kugo.SpotOrdersRequest{
		ClientOid:   "123",
		Side:        "buy",
		Symbol:      "BTC-USDT",
		Type:        "limit",
		TradeType:   "TRADE",
		Price:       decimal.NewFromFloat(450000),
		Size:        decimal.NewFromFloat(0.00001),
		TimeInForce: "IOC",
	}
	result, err := instance.SpotOrder(req)
	t.Log(result, err)
}

func TestSpotOrderFills(t *testing.T) {
	req := &kugo.SpotOrderFillsRequest{
		//OrderId:   "6424f39e926d4e0001c14029",
		TradeType: "TRADE",
	}
	result, err := instance.SpotOrderFills(req, 1, 50)
	t.Log(result, err)
}

func TestFutureAccount(t *testing.T) {
	instance.Set(kugo.SetApiKey(futureAccessKey, futureSecretKey, futurePassphrase))
	t.Log(instance)
	account, err := instance.FutureAccount("USDT")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", account)
}

func TestFutureOrder(t *testing.T) {
	instance.Set(kugo.SetApiKey(futureAccessKey, futureSecretKey, futurePassphrase))
	req := &kugo.FutureOrderRequest{
		ClientOid:   "123",
		Side:        "buy",
		Symbol:      "ETH-USDT",
		Type:        "limit",
		Price:       decimal.NewFromFloat(40),
		Size:        1,
		TimeInForce: "IOC",
	}
	result, err := instance.FutureOrder(req)
	t.Log(result, err)
}
