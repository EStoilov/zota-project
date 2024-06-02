package services

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"zota-project/models"
)

func TestCreateSignature(t *testing.T) {
	ds := NewDepositService("test_secret_key")
	request := &models.DepositRequest{
		EndpointID:      "test_endpoint",
		MerchantOrderID: "test_order",
		OrderAmount:     100.0,
		CustomerEmail:   "test@example.com",
	}

	ds.CreateSignature(request)

	expectedSignature := fmt.Sprintf("%x", sha256.Sum256([]byte("test_endpointtest_order100.00test@example.comtest_secret_key")))

	if request.Signature != expectedSignature {
		t.Errorf("expected signature to be %s, got %s", expectedSignature, request.Signature)
	}
}