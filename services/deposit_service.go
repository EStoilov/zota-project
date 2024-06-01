package services

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"zota-project/models"
)

type DepositService struct {
	SecretKey string
}

func NewDepositService(secretKey string) *DepositService {
	return &DepositService{SecretKey: secretKey}
}

func (ds *DepositService) CreateSignature(request *models.DepositRequest) {
	signatureString := fmt.Sprintf("%s%s%.2f%s%s",
		request.EndpointID,
		request.MerchantOrderID,
		request.OrderAmount,
		request.CustomerEmail,
		ds.SecretKey)
	hash := sha256.New()
	hash.Write([]byte(signatureString))
	request.Signature = fmt.Sprintf("%x", hash.Sum(nil))
}

func (ds *DepositService) SendDepositRequest(request models.DepositRequest) (models.DepositResponse, error) {
	url := "https://api.zotapay.com/deposit"
	response, err := utils.SendPostRequest(url, request)
	if err != nil {
		return models.DepositResponse{}, err
	}

	var depositResponse models.DepositResponse
	if err := json.NewDecoder(response.Body).Decode(&depositResponse); err != nil {
		return models.DepositResponse{}, err
	}

	return depositResponse, nil
}