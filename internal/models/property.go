package models

import (
	"time"
)

type Property struct {
	ID          string    `json:"id"`
	Owner       string    `json:"owner"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
