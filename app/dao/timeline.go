package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

var _ repository.Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (tl *timeline) GetStatusTimeline(ctx context.Context, statusID int, limitNum int) (*object.Timeline, error) {
	timeline := new(object.Timeline)
	rows, err := tl.db.QueryxContext(ctx, "SELECT * FROM timeline WHERE id >= ? LIMIT ?", statusID, limitNum)
	if err != nil {
		return nil, fmt.Errorf("failed to query timeline from db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(timeline); err != nil {
			return nil, fmt.Errorf("failed to scan status: %w", err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during rows iteration: %w", err)
	}
	return timeline, nil
}
