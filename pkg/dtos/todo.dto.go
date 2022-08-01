package dtos

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/datatypes"
)

type Assignee struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	ImageUrl string `json:"imageUrl"`
}

type AssigneeSlice []Assignee

// user - DTOs
type CreateTodo struct {
	ObjectId       string         `json:"objectId"`
	TodoStatus     string         `json:"todoStatus"` // open, close, completed
	ActionText     string         `json:"actionText"`
	Description    string         `json:"description"`
	ListedBy       string         `json:"listedBy"`
	ExpirationTime time.Time      `json:"expiryTime"`
	ExpirationDate datatypes.Date `json:"expiryDate"`
	Assignee       *AssigneeSlice `json:"assignee"`
	// UserId         uint           `json:"user"`
}

// Customs Valuer/Scanner for AssigneeSlice
func (asg AssigneeSlice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(asg)
	return string(bytes), err
}

func (asg *AssigneeSlice) Scan(input interface{}) error {
	switch value := input.(type) {
	case string:
		return json.Unmarshal([]byte(value), asg)
	case []byte:
		return json.Unmarshal(value, asg)
	default:
		return errors.New("not supported")
	}
}
