package services

import (
	"crypto/sha256"
	"fmt"
	"zota-project/models"
)

type DepositService struct {
	secretKey string
}

func NewDepositService(secretKey string) *DepositService {
	return &DepositService{secretKey: secretKey}
}

func (ds *DepositService) CreateSignature(request *models.DepositRequest) {
	data := fmt.Sprintf("%s%s%.2f%s%s", request.EndpointID, request.MerchantOrderID, request.OrderAmount, request.CustomerEmail, ds.secretKey)
	request.Signature = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}