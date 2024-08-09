package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	statusUsecase usecase.StatusUsecase
}

// Create Handler for `/v1/statuses/`
func NewRouter(ar repository.Account, us usecase.StatusUsecase) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		statusUsecase: us,
	}
	r.Get("/timeline/public", h.GetTimeline)

	return r
}
