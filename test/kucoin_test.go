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

func TestSpotSymbols(t *testing.T) {
	instance.Set(kugo.SetApiKey(accessKey, secretKey, passphrase))
	symbols, err := instance.SpotSymbols("USDS")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", symbols)
}

func TestAccounts(t *testing.T) {
	instance.Set(kugo.SetApiKey(accessKey, secretKey, passphrase))
	accounts, err := instance.SpotAccount("BTC", "trade")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %+v", accounts)
}

func TestSpotOrder(t *testing.T) {
	instance.Set(kugo.SetApiKey(accessKey, secretKey, passphrase))
	req := &kugo.SpotOrdersRequest{
		ClientOid:   "123",
		Side:        "buy",
		Symbol:      "BTC-USDT",
		Type:        "limit",
		TradeType:   "TRADE",
		Price:       decimal.NewFromFloat(10000),
		Size:        decimal.NewFromFloat(0.00001),
		TimeInForce: "GTC",
	}
	result, err := instance.SpotOrder(req)
	t.Log(result, err)
}

func TestMarginSpotOrder(t *testing.T) {
	instance.Set(kugo.SetApiKey(accessKey, secretKey, passphrase))
	req := &kugo.SpotMarginOrderRequest{
		ClientOid:   "123",
		Side:        "buy",
		Symbol:      "BTC-USDT",
		Type:        "limit",
		Price:       decimal.NewFromFloat(10000),
		Size:        decimal.NewFromFloat(0.00001),
		TimeInForce: "GTC",
	}
	result, err := instance.SpotMarginOrder(req)
	t.Log(result, err)
}

func TestSpotOrderFills(t *testing.T) {
	instance.Set(kugo.SetApiKey(accessKey, secretKey, passphrase))
	req := &kugo.SpotOrderFillsRequest{
		//OrderId:   "6424f39e926d4e0001c14029",
		TradeType: "TRADE",
	}
	result, err := instance.SpotOrderFills(req, 1, 50)
	t.Log(result, err)
}

func TestSpotOrderList(t *testing.T) {
	instance.Set(kugo.SetApiKey(accessKey, secretKey, passphrase))
	req := &kugo.SpotOrderListRequest{
		Status:    "",
		Symbol:    "",
		Side:      "",
		Type:      "",
		TradeType: "TRADE",
		StartAt:   0,
		EndAt:     0,
	}
	result, err := instance.SpotOrderList(req, 1, 10)
	t.Log(result, err)
}

func TestSpotOrderOne(t *testing.T) {
	instance.Set(kugo.SetApiKey(accessKey, secretKey, passphrase))
	result, err := instance.SpotOrderOne("642a8cfa926d4e0001c86207")
	t.Log(result, err)
}

func TestSpotOrderCancel(t *testing.T) {
	instance.Set(kugo.SetApiKey(accessKey, secretKey, passphrase))
	result, err := instance.SpotOrderCancel("642a8cfa926d4e0001c86207")
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
		Symbol:      "ETHUSDTM",
		Leverage:    decimal.NewFromFloat(1),
		Type:        "limit",
		Price:       decimal.NewFromFloat(1800),
		Size:        1,
		TimeInForce: "GTC",
	}
	result, err := instance.FutureOrder(req)
	t.Log(result, err)
}

func TestFutureOrderCancel(t *testing.T) {
	instance.Set(kugo.SetApiKey(futureAccessKey, futureSecretKey, futurePassphrase))
	result, err := instance.FutureOrderCancel("642b807bcac00c0001177123")
	t.Log(result, err)
}

func TestFutureOrderList(t *testing.T) {
	instance.Set(kugo.SetApiKey(futureAccessKey, futureSecretKey, futurePassphrase))
	req := &kugo.FutureOrderListRequest{
		Status:  "",
		Symbol:  "",
		Side:    "",
		Type:    "",
		StartAt: 0,
		EndAt:   0,
	}
	result, err := instance.FutureOrderList(req, 1, 10)
	t.Log(result, err)
}

func TestFutureOrderOne(t *testing.T) {
	instance.Set(kugo.SetApiKey(futureAccessKey, futureSecretKey, futurePassphrase))
	result, err := instance.FutureOrderOne("642b807bcac00c0001177123")
	t.Log(result, err)
}

func TestFutureOrderFills(t *testing.T) {
	instance.Set(kugo.SetApiKey(futureAccessKey, futureSecretKey, futurePassphrase))
	req := &kugo.FutureOrderFillsRequest{
		//OrderId:   "6424f39e926d4e0001c14029",
	}
	result, err := instance.FutureOrderFills(req, 1, 50)
	t.Log(result, err)
}

func TestFuturePosition(t *testing.T) {
	instance.Set(kugo.SetApiKey(futureAccessKey, futureSecretKey, futurePassphrase))
	result, err := instance.FuturePosition("XBTUSDTM")
	t.Log(result, err)
}

func TestFutureSymbols(t *testing.T) {
	instance.Set(kugo.SetApiKey(futureAccessKey, futureSecretKey, futurePassphrase))
	result, err := instance.FutureSymbols()
	t.Log(result, err)
}
