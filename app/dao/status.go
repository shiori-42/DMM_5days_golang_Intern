package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

func (s *status) FindByStatusID(ctx context.Context, statusID int) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", statusID).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}

func (s *status) CreateStatus(ctx context.Context, tx *sqlx.Tx, acc *object.Account, status *object.Status) error {
	_, err := tx.Exec("insert into status (account_id, content, created_at) values (?, ?, ?)",
		acc.ID, status.Content, status.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to insert account: %w", err)
	}

	return nil
}
