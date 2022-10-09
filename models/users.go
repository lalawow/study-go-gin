package models

import (
	"time"
)

type User struct {
	Id        int        `json:"id"`
	Data      string     `json:"data" binding:"required"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}

// TableName is Database Table Name of this model
func (e *User) Users() string {
	return "users"
}
