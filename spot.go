package kugo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// SpotOrder POST /api/v1/orders
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
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// SpotMarginOrder POST /api/v1/margin/order
func (kc *Kucoin) SpotMarginOrder(req *SpotMarginOrderRequest) (*SpotMarginOrderData, error) {
	uri := UriSpotMarginOrder
	p, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := kc.do(kc.spotEndpoint, http.MethodPost, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &SpotMarginOrderResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// SpotOrderFills GET /api/v1/fills
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
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// SpotOrderCancel DELETE /api/v1/orders/{orderId}
func (kc *Kucoin) SpotOrderCancel(orderId string) (*SpotOrderCancelData, error) {
	uri := fmt.Sprintf(UriSpotOrderCancel, orderId)
	resp, err := kc.do(kc.spotEndpoint, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	respStruct := &SpotOrderCancelResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// SpotOrderList GET /api/v1/orders
func (kc *Kucoin) SpotOrderList(req *SpotOrderListRequest, currentPage, pageSize int) (*SpotOrderListData, error) {
	uri := UriSpotOrders
	p := map[string]string{}
	p["currentPage"] = strconv.Itoa(currentPage)
	p["pageSize"] = strconv.Itoa(pageSize)
	if len(req.Status) != 0 {
		p["status"] = req.Status
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
	if len(req.TradeType) != 0 {
		p["tradeType"] = req.TradeType
	}
	if req.StartAt != 0 {
		p["startAt"] = strconv.Itoa(int(req.StartAt))
	}
	if req.EndAt != 0 {
		p["endAt"] = strconv.Itoa(int(req.EndAt))
	}

	resp, err := kc.do(kc.spotEndpoint, http.MethodGet, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &SpotOrderListResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// SpotOrderOne GET /api/v1/orders/{orderId}
func (kc *Kucoin) SpotOrderOne(orderId string) (*SpotOrderOneData, error) {
	uri := fmt.Sprintf(UriSpotOrderOne, orderId)
	resp, err := kc.do(kc.spotEndpoint, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	respStruct := &SpotOrderOneResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg)!=0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}
