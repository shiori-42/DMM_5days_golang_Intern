package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	GetStatusTimeline(ctx context.Context, statusID int, limitNum int) (*object.Timeline, error)
}
