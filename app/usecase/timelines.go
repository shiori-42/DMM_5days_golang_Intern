package usecase

import (
	"context"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

func NewTimeline(db *sqlx.DB, timlineRepo repository.Timeline) *timline {
	return &timline{
		db:          db,
		timlineRepo: timlineRepo,
	}
}

type Timelineecase interface {
	GetStatusTimeline(ctx context.Context, statusIDStr string) (*GetStatusDTO, error)
}

type timeline struct {
	db           *sqlx.DB
	TimelineRepo repository.Status
}

type GetTimelineDTO struct {
	Timeline *object.Timeline
	Account  *object.Account
}

var _ TimelineUsecase = (*timeline)(nil)

func (tl *timeline) GetStatusTimeline(ctx context.Context, statusIDStr string) (*GetTimelineDTO, error) {

	statusID, err := strconv.Atoi(statusIDStr)
	if err != nil {
		return nil, err
	}
	timeline, err := GetStatusTimeline(ctx, statusID, limitNum)
	if err != nil {
		return nil, err
	}
	return &GetTimelineDTO{
		Timeline: timeline,
	}, nil

}
