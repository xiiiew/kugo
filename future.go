package kugo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// FutureAccount GET /api/v1/account-overview
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
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// FutureOrder POST /api/v1/orders
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
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// FutureOrderCancel DELETE /api/v1/orders/{orderId}
func (kc *Kucoin) FutureOrderCancel(orderId string) (*FutureOrderCancelData, error) {
	uri := fmt.Sprintf(UriFutureOrderCancel, orderId)
	resp, err := kc.do(kc.futureEndpoint, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	respStruct := &FutureOrderCancelResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// FutureOrderList GET /api/v1/orders
func (kc *Kucoin) FutureOrderList(req *FutureOrderListRequest, currentPage, pageSize int) (*FutureOrderListData, error) {
	uri := UriFutureOrders
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
	if req.StartAt != 0 {
		p["startAt"] = strconv.Itoa(int(req.StartAt))
	}
	if req.EndAt != 0 {
		p["endAt"] = strconv.Itoa(int(req.EndAt))
	}

	resp, err := kc.do(kc.futureEndpoint, http.MethodGet, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &FutureOrderListResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// FutureOrderOne GET /api/v1/orders/{orderId}
func (kc *Kucoin) FutureOrderOne(orderId string) (*FutureOrderOneData, error) {
	uri := fmt.Sprintf(UriFutureOrderOne, orderId)
	resp, err := kc.do(kc.futureEndpoint, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	respStruct := &FutureOrderOneResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// FutureOrderFills GET /api/v1/fills
func (kc *Kucoin) FutureOrderFills(req *FutureOrderFillsRequest, currentPage, pageSize int) (*FutureOrderFillsData, error) {
	uri := UriFutureOrderFills
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

	resp, err := kc.do(kc.futureEndpoint, http.MethodGet, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &FutureOrderFillsResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}

// FuturePosition GET /api/v1/position
func (kc *Kucoin) FuturePosition(symbol string) (*FuturePositionData, error) {
	uri := UriFuturePosition
	p := make(map[string]string, 0)
	if len(symbol) != 0 {
		p["symbol"] = symbol
	}
	resp, err := kc.do(kc.futureEndpoint, http.MethodGet, uri, p)
	if err != nil {
		return nil, err
	}

	respStruct := &FuturePositionResponse{}
	if err = json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return nil, err
	}
	if respStruct.Code != "200000" && respStruct.Code != "200" || len(respStruct.Msg) != 0 {
		return nil, errors.New(respStruct.Msg)
	}
	return &respStruct.Data, nil
}
