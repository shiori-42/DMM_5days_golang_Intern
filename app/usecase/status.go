package usecase

import (
	"context"
	"fmt"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type StatusUsecase interface {
	CreateStatus(ctx context.Context, content string, account *object.Account) (*CreateStatusDTO, error)
	GetStatus(ctx context.Context, statusIDStr string) (*GetStatusDTO, error)
	GetAccountByAccountID(ctx context.Context, accountID int) (*GetStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type CreateStatusDTO struct {
	Status *object.Status
}

type GetStatusDTO struct {
	Status  *object.Status
	Account *object.Account
}

var _ StatusUsecase = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) CreateStatus(ctx context.Context, content string, account *object.Account) (*CreateStatusDTO, error) {
	status := object.NewStatus(content, int(account.ID))

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if err := s.statusRepo.CreateStatus(ctx, tx, account, status); err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Status: status,
	}, nil
}

func (a *status) GetStatus(ctx context.Context, statusIDStr string) (*GetStatusDTO, error) {

	statusID, err := strconv.Atoi(statusIDStr)
	if err != nil {
		return nil, err
	}
	status, err := a.statusRepo.FindByStatusID(ctx, statusID)
	if err != nil {
		return nil, err
	}
	return &GetStatusDTO{
		Status: status,
	}, nil

}

func (s *status) GetAccountByAccountID(ctx context.Context, accountID int) (*GetStatusDTO, error) {
	account, err := s.statusRepo.FindAccountByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n",account)
	return &GetStatusDTO{
		Account: account,
	}, nil

}
