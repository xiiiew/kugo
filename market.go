package kugo

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Symbols GET /api/v2/symbols
func (kc *Kucoin) Symbols(market string) ([]SymbolsData, error) {
	uri := UriSymbols
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
