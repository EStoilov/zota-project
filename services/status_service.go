package services

import (
	"crypto/sha256"
	"fmt"
	"zota-project/models"
)

type StatusService struct {
	secretKey string
}

func NewStatusService(secretKey string) *StatusService {
	return &StatusService{secretKey: secretKey}
}

func (ss *StatusService) CreateSignature(request *models.StatusRequest) {
	data := fmt.Sprintf("%s%s%s", request.MerchantOrderID, request.OrderID, ss.secretKey)
	request.Signature = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}