package http

import (
	"encoding/json"
	"net/http"
	"zota-project/models"
	"zota-project/services"
	"zota-project/utils"
)

type Handler struct {
	depositService *services.DepositService
	statusService  *services.StatusService
}

func NewHandler(depositService *services.DepositService, statusService *services.StatusService) *Handler {
	return &Handler{
		depositService: depositService,
		statusService:  statusService,
	}
}

func (h *Handler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	var request models.DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.depositService.CreateSignature(&request)

	response := models.DepositResponse{
		Code:    "",
		Message: "",
		Data: models.DepositData{
			DepositURL:     "",
			MerchantOrderID: request.MerchantOrderID,
			OrderID:        "",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not encode response")
		return
	}
}

func (h *Handler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	var request models.StatusRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.statusService.CreateSignature(&request)

	response := models.StatusResponse{
		Code:    "",
		Message: "",
		Data: models.StatusData{
			Status:         "",
			MerchantOrderID: request.MerchantOrderID,
			OrderID:        request.OrderID,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not encode response")
		return
	}
}

func (h *Handler) PollStatusHandler(w http.ResponseWriter, r *http.Request) {
	var request models.StatusRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	h.statusService.CreateSignature(&request)

	response := models.StatusResponse{
		Code:    "",
		Message: "",
		Data: models.StatusData{
			Status:         "",
			MerchantOrderID: request.MerchantOrderID,
			OrderID:        request.OrderID,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not encode response")
		return
	}
}
