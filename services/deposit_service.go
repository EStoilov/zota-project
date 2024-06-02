package services

import (
	"crypto/sha256"
	"fmt"
	"zota-project/models"
	"zota-project/config"
)

type DepositService struct {
	config *config.Config
}

func NewDepositService(cfg *config.Config) *DepositService {
	return &DepositService{config: cfg}

}

func (ds *DepositService) CreateSignature(request *models.DepositRequest) {
	request.OrderCurrency = ds.config.Currency
	request.CustomerFirstName = "John"
	request.CustomerPhone = "+1234567890" 
	request.CustomerIP = "192.168.1.1" 
	request.RedirectURL = "https://example.com/redirect"
	request.CheckoutURL = "https://example.com/checkout"
	request.Language = "en" 

	data := fmt.Sprintf("%s%s%.2f%s%s", request.EndpointID, request.MerchantOrderID, request.OrderAmount, request.CustomerEmail, ds.config.SecretKey)
	request.Signature = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}