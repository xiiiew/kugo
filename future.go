package kugo

import (
	"encoding/json"
	"errors"
	"net/http"
)

// FutureAccount /api/v1/account-overview
func (kc *Kucoin) FutureAccount(currency string) (*FutureAccountData, error) {
	uri := UriFutureAccount
	p := map[string]string{}
	if len(currency) != 0 {
		p["currency"] = currency
	}

	resp, err := kc.do(kc.futureEndpoint, http.MethodGet, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &FutureAccountResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// FutureOrder /api/v1/orders
func (kc *Kucoin) FutureOrder(req *FutureOrderRequest) (*FutureOrderData, error) {
	uri := UriFutureOrders
	p, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := kc.do(kc.futureEndpoint, http.MethodPost, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &FutureOrderResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}
