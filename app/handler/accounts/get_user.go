package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handle request for `POST /v1/accounts`
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	ctx := r.Context()

	dto, err := h.accountUsecase.GetUser(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
