package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zota-project/models"
	"zota-project/services"
)

func TestDepositHandler(t *testing.T) {
	depositService := services.NewDepositService("5f4a6fcf-9048-4a0b-afc2-ed92d60fb1bf")
	statusService := services.NewStatusService("5f4a6fcf-9048-4a0b-afc2-ed92d60fb1bf")
	handler := NewHandler(depositService, statusService)

	depositRequest := models.DepositRequest{
		EndpointID:      "402334",
		MerchantOrderID: "BUGBOUNTY231",
		OrderAmount:     100.0,
		CustomerEmail:   "test@example.com",
	}

	body, _ := json.Marshal(depositRequest)
	req, err := http.NewRequest("POST", "/deposit", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.DepositHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"code":"","message":"","data":{"depositUrl":"","merchantOrderID":"","orderID":""}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
