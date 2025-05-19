package handler

import (
	"encoding/json"
	"net/http"
	"noisely/internal/usecase"
)

type SchemaHandler struct {
	uc *usecase.SchemaUsecase
}

func NewSchemaHandler(uc *usecase.SchemaUsecase) *SchemaHandler {
	return &SchemaHandler{uc: uc}
}

func (h *SchemaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: преобразовать req в entities.Schema и вызвать usecase
	w.WriteHeader(http.StatusCreated)
}
