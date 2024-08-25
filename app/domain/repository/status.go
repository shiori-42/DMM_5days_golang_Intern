package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	CreateStatus(ctx context.Context, tx *sqlx.Tx, acc *object.Account, status *object.Status) error
	FindByStatusID(ctx context.Context, statusID int) (*object.Status, error)
	FindAccountByAccountID(ctx context.Context, AccountID int) (*object.Account, error)
}
