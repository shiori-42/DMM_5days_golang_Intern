package accounts

// import (
// 	"encoding/json"
// 	"net/http"
// )

// // Request body for `POST /v1/accounts/update_credentials`
// type UpgradeRequest struct {
// 	DisplayName string
// 	Note string
// 	Avater string
// 	Header string
// }


// // Handle request for `POST /v1/accounts/update_credentials`
// func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
// 	var req UpgradeRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	ctx := r.Context()

// 	dto, err := h.accountUsecase.Update(ctx, req.Username, req.Password)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(dto.Account); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
