package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// type MediaAttachment struct {
//     ID          int    `json:"id"`
//     Type        string `json:"type"`
//     URL         string `json:"url"`
//     Description string `json:"description"`
// }

type Account struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	DisplayName *string   `json:"display_name"`
	CreateAt    time.Time `json:"create_at"`
	// FollowersCount int       `json:"followers_count"`
	// FollowingCount int       `json:"following_count"`
	Note   *string `json:"note"`
	Avatar *string `json:"avatar"`
	Header *string `json:"header"`
}

type Response struct {
	ID       int       `json:"id"`
	Account  Account   `json:"account"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
	// MediaAttachments []MediaAttachment `json:"media_attachments"`
}

// Handle request for `POST /v1/statuses`
func (h *handler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	statusIDStr := chi.URLParam(r, "id")

	ctx := r.Context()

	status, err := h.statusUsecase.GetTimeline(ctx, statusIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	account, err := h.statusUsecase.GetAccountByAccountID(ctx, status.Status.AccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(status.Status.ID)
	res := &Response{
		ID: status.Status.ID,
		Account: Account{
			ID:          account.Account.ID,
			Username:    account.Account.Username,
			DisplayName: account.Account.DisplayName,
			CreateAt:    account.Account.CreateAt,
			Note:        account.Account.Note,
			Avatar:      account.Account.Avatar,
			Header:      account.Account.Header,
		},
		Content:  status.Status.Content,
		CreateAt: status.Status.CreatedAt,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
