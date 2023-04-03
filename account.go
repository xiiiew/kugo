package kugo

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Accounts GET /api/v1/account
func (kc *Kucoin) Accounts(currency, _type string) ([]AccountsData, error) {
	uri := UriAccounts
	p := map[string]string{}
	if len(currency) != 0 {
		p["currency"] = currency
	}
	if len(_type) != 0 {
		p["type"] = _type
	}

	resp, err := kc.do(kc.spotEndpoint, http.MethodGet, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &AccountsResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return respStruct.Data, nil
}
