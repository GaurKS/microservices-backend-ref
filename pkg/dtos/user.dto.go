package dtos

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/datatypes"
)

type Education struct {
	Institute string `json:"institute"`
	Course    string `json:"course"`
	Grade     string `json:"grade"`
	Code      string `json:"code"`
}

type WorkExperience struct {
	Organisation string         `json:"organisation"`
	Role         string         `json:"role"`
	StartDate    datatypes.Date `json:"startDate"`
	EndDate      datatypes.Date `json:"endDate"`
}

type EducationSlice []Education
type WorkExperienceSlice []WorkExperience

// user - DTOs
type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditUser struct {
	Name           string               `json:"name,omitempty"`
	Age            int                  `json:"age,omitempty"`
	Gender         string               `json:"gender,omitempty"` // Male, Female, Other
	Role           string               `json:"role,omitempty"`   // CA, PA, SA
	ImageUrl       string               `json:"imageUrl,omitempty"`
	Email          string               `json:"email,omitempty"`
	Interest       datatypes.JSON       `json:"interest,omitempty"`
	Social         datatypes.JSON       `json:"social,omitempty"`
	Education      *EducationSlice      `json:"education,omitempty"`
	WorkExperience *WorkExperienceSlice `json:"workExperience,omitempty"`
}

func (edu EducationSlice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(edu)
	return string(bytes), err
}

func (edu *EducationSlice) Scan(input interface{}) error {
	switch value := input.(type) {
	case string:
		return json.Unmarshal([]byte(value), edu)
	case []byte:
		return json.Unmarshal(value, edu)
	default:
		return errors.New("not supported")
	}
}

// Customs Valuer/Scanner for WorkExperienceSlice
func (we WorkExperienceSlice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(we)
	return string(bytes), err
}

func (we *WorkExperienceSlice) Scan(input interface{}) error {
	switch value := input.(type) {
	case string:
		return json.Unmarshal([]byte(value), we)
	case []byte:
		return json.Unmarshal(value, we)
	default:
		return errors.New("not supported")
	}
}
