package http

import (
	"encoding/json"
	"net/http"
	"zota-project/models"
	"zota-project/services"
)

type Handler struct {
	DepositService *services.DepositService
	StatusService  *services.StatusService
}

func NewHandler(ds *services.DepositService, ss *services.StatusService) *Handler {
	return &Handler{DepositService: ds, StatusService: ss}
}

func (h *Handler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	var request models.DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.DepositService.CreateSignature(&request)

	response, err := h.DepositService.SendDepositRequest(request)
	if err != nil {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	var request models.StatusRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.StatusService.CreateSignature(&request)

	response, err := h.StatusService.CheckOrderStatus(request)
	if err != nil {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) PollStatusHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		MerchantOrderID string `json:"merchantOrderID"`
		OrderID         string `json:"orderID"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := h.StatusService.PollOrderStatus(request.MerchantOrderID, request.OrderID)
	if err != nil {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}