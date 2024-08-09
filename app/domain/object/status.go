package object

import (
	"time"
)

type Status struct {
	ID        int            `json:"id,omitempty"`
	AccountID int            `json:"account_id,omitempty" db:"account_id"`
	URL       *string        `json:"url,omitempty" db:"url"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at,omitempty" db:"created_at"`
}

func NewStatus(content string, accountID int) *Status {
	return &Status{
		AccountID: accountID,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
