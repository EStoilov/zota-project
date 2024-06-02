package models

type StatusRequest struct {
	MerchantOrderID string `json:"merchantOrderID"`
	OrderID         string `json:"orderID"`
	Signature       string `json:"signature"`
}

type StatusResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    StatusData  `json:"data,omitempty"`
}

type StatusData struct {
	MerchantOrderID   string      `json:"merchantOrderID"`
	OrderID           string      `json:"orderID"`
	Status            string      `json:"status"`
	OriginalAmount    float64     `json:"originalAmount,omitempty"`
	AmountChanged     float64     `json:"amountChanged,omitempty"`
	AmountRounded     float64     `json:"amountRounded,omitempty"`
	AmountManipulated float64     `json:"amountManipulated,omitempty"`
	ExtraData         interface{} `json:"extraData,omitempty"`
}