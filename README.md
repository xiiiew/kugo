# KuGo

Go SDK for KuCoin API

> This SDK is currently available for KuCoin V2 API KEY. All API request parameters and response details are described in the documentation at [https://docs.kucoin.com](https://docs.kucoin.com).

## Install

```shell
go get github.com/xiiiew/kugo
```

## REST API Support

<details>
<summary>Spot</summary>

|             URI               |
|-------------------------------|
| /api/v2/symbols               |
| /api/v1/accounts              |
| /api/v1/orders                |
| /api/v1/fills                 |

</details>

<details>
<summary>Future</summary>

|             URI               |
|-------------------------------|
| /api/v1/account-overview      |
| /api/v1/orders                |

</details>

## Usage

### Create Instance

```golang
// Default instance
instance, err := kugo.NewKucoin()

// Set Kucoin V2 API Key
instance, err := kugo.NewKucoin(
    kugo.SetApiKey("accessKey", "secretKey", "passphrase"),
)

// Set environment
instance, err := kugo.NewKucoin(
    kugo.SetSpotEndpoint("https://openapi-sandbox.kucoin.com"),
    kugo.SetFutureEndpoint("https://api-sandbox-futures.kucoin.com"),
)

// Debug mode. Default output of debug information to the console
instance, err := kugo.NewKucoin(
    kugo.SetDebug(true),
)

// Set the output mode of debug information (e.g. to log files)
instance, err := kugo.NewKucoin(
    kugo.SetDebug(true),
    kugo.SetRequestLog(func(i ...interface{}) {
        // Output request log
    }),
    kugo.SetResponseLog(func(i ...interface{}) {
        // Output response log
    }),
)

// Set HTTP client
uProxy, _ := url.Parse("http://127.0.0.1:7890")
instance, err := kugo.NewKucoin(
    kugo.SetClient(&http.Client{Transport: &http.Transport{
			Proxy:             http.ProxyURL(uProxy),
			DisableKeepAlives: true},
			Timeout: 10 * time.Second},
    ),
),

```

### Examples

> See the test case for more examples.

```golang
// Spot symbols
symbols, err := instance.Symbols("USDS")
if err != nil {
    t.Fatal(err)
}
t.Logf("result: %+v", symbols)

// Spot order
req := &kugo.SpotOrdersRequest{
    ClientOid:   "123",
    Side:        "buy",
    Symbol:      "BTC-USDT",
    Type:        "limit",
    TradeType:   "TRADE",
    Price:       decimal.NewFromFloat(450000),
    Size:        decimal.NewFromFloat(1),
    TimeInForce: "IOC",
}
result, err := instance.SpotOrder(req)
t.Log(result, err)
```
