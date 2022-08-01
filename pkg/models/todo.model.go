package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Assignee struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	ImageUrl string `json:"imageUrl"`
}

type AssigneeSlice []Assignee

type Todo struct {
	gorm.Model     `json:"-"`
	ObjectId       string         `json:"objectId,omitempty"`
	TodoStatus     string         `json:"todoStatus,omitempty"` // open, close, completed
	ActionText     string         `json:"actionText,omitempty"`
	Description    string         `json:"description,omitempty"`
	ListedBy       string         `json:"-"`
	UserId         string         `gorm:"" json:"-"`
	ExpirationTime time.Time      `json:"expiryTime,omitempty"`
	ExpirationDate datatypes.Date `json:"expiryDate,omitempty"`
	Assignee       AssigneeSlice  `json:"assignee,omitempty"`
}
