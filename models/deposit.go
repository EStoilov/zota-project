package models

type DepositRequest struct {
	EndpointID                string  `json:"endpointID"`
	MerchantOrderID           string  `json:"merchantOrderID"`
	OrderAmount               float64 `json:"orderAmount"`
	OrderCurrency             string  `json:"orderCurrency"`
	CustomerEmail             string  `json:"customerEmail"`
	CustomerFirstName         string  `json:"customerFirstName"`
	CustomerLastName          string  `json:"customerLastName"`
	CustomerPhone             string  `json:"customerPhone"`
	CustomerIP                string  `json:"customerIP"`
	CustomerBankCode          string  `json:"customerBankCode,omitempty"`
	CustomerBankAccountNumber string  `json:"customerBankAccountNumber,omitempty"`
	RedirectURL               string  `json:"redirectUrl"`
	CallbackURL               string  `json:"callbackUrl,omitempty"`
	CheckoutURL               string  `json:"checkoutUrl"`
	CustomParam               string  `json:"customParam,omitempty"`
	Language                  string  `json:"language,omitempty"`
	Signature                 string  `json:"signature"`
}

type DepositResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    DepositData `json:"data,omitempty"`
}

type DepositData struct {
	DepositURL      string `json:"depositUrl"`
	MerchantOrderID string `json:"merchantOrderID"`
	OrderID         string `json:"orderID"`
}