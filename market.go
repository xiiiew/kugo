package kugo

import (
	"encoding/json"
	"errors"
	"net/http"
)

// SpotSymbols GET /api/v2/symbols
func (kc *Kucoin) SpotSymbols(market string) ([]SymbolsData, error) {
	uri := UriSpotSymbols
	p := map[string]string{}
	if len(market) != 0 {
		p["market"] = market
	}

	resp, err := kc.do(kc.spotEndpoint, http.MethodGet, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &SymbolsResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return respStruct.Data, nil
}

// FutureSymbols GET /api/v1/contracts/active
func (kc *Kucoin) FutureSymbols() ([]FutureSymbolData, error) {
	uri := UriFutureSymbols

	resp, err := kc.do(kc.futureEndpoint, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	respStruct := &FutureSymbolResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return respStruct.Data, nil
}
