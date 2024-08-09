package usecase

import (
	"context"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type AccountUsecase interface {
	Create(ctx context.Context, username, password string) (*CreateAccountDTO, error)
	GetUser(ctx context.Context, username string) (*FindAccountByAccountIDDTO, error)
}

type account struct {
	db          *sqlx.DB
	accountRepo repository.Account
}

type CreateAccountDTO struct {
	Account *object.Account
}

type FindAccountByAccountIDDTO struct {
	Account *object.Account
}

var _ AccountUsecase = (*account)(nil)

func NewAcocunt(db *sqlx.DB, accountRepo repository.Account) *account {
	return &account{
		db:          db,
		accountRepo: accountRepo,
	}
}

func (a *account) Create(ctx context.Context, username, password string) (*CreateAccountDTO, error) {
	acc, err := object.NewAccount(username, password)
	if err != nil {
		return nil, err
	}

	tx, err := a.db.Beginx()
	if err != nil {
		return nil, err
	}

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		tx.Rollback()
	// 	}

	// 	tx.Commit()
	// }()

	if err := a.accountRepo.Create(ctx, tx, acc); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Print(err)
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &CreateAccountDTO{
		Account: acc,
	}, nil
}

func (a *account) GetUser(ctx context.Context, username string) (*FindAccountByAccountIDDTO, error) {

	account, err := a.accountRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &FindAccountByAccountIDDTO{
		Account: account,
	}, nil

}
