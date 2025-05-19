package handler

import (
	"encoding/json"
	"net/http"
	"noisely/internal/usecase"
)

type DatasetHandler struct {
	uc *usecase.DatasetUsecase
}

func NewDatasetHandler(uc *usecase.DatasetUsecase) *DatasetHandler {
	return &DatasetHandler{uc: uc}
}

func (h *DatasetHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: преобразовать req в entities.Dataset и вызвать usecase
	w.WriteHeader(http.StatusCreated)
}
