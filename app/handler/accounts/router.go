package accounts

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	accountUsecase usecase.AccountUsecase
}

// Create Handler for `/v1/accounts/`
func NewRouter(u usecase.AccountUsecase) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		accountUsecase: u,
	}
	r.Post("/", h.Create)
	r.Get("/{username}", h.GetUser)

	return r
}
