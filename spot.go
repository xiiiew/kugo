package kugo

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// SpotOrder /api/v1/orders
func (kc *Kucoin) SpotOrder(req *SpotOrdersRequest) (*SpotOrderData, error) {
	uri := UriSpotOrders
	p, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := kc.do(kc.spotEndpoint, http.MethodPost, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &SpotOrderResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// SpotOrderFills /api/v1/fills
func (kc *Kucoin) SpotOrderFills(req *SpotOrderFillsRequest, currentPage, pageSize int) (*SpotOrderFillsData, error) {
	uri := UriSpotOrderFills
	p := map[string]string{}
	p["currentPage"] = strconv.Itoa(currentPage)
	p["pageSize"] = strconv.Itoa(pageSize)
	if len(req.OrderId) != 0 {
		p["orderId"] = req.OrderId
	}
	if len(req.Symbol) != 0 {
		p["symbol"] = req.Symbol
	}
	if len(req.Side) != 0 {
		p["side"] = req.Side
	}
	if len(req.Type) != 0 {
		p["type"] = req.Type
	}
	if req.StartAt != 0 {
		p["startAt"] = strconv.Itoa(int(req.StartAt))
	}
	if req.EndAt != 0 {
		p["endAt"] = strconv.Itoa(int(req.EndAt))
	}
	if len(req.OrderId) != 0 {
		p["tradeType"] = req.TradeType
	}

	resp, err := kc.do(kc.spotEndpoint, http.MethodGet, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &SpotOrderFillsResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}
