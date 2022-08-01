package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Customs Valuer/Scanner for EducationSlice
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

// Customs Valuer/Scanner for AssigneeSlice
func (asgn AssigneeSlice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(asgn)
	return string(bytes), err
}

func (asgn *AssigneeSlice) Scan(input interface{}) error {
	switch value := input.(type) {
	case string:
		return json.Unmarshal([]byte(value), asgn)
	case []byte:
		return json.Unmarshal(value, asgn)
	default:
		return errors.New("not supported")
	}
}

func (user *User) CheckPassword(pswd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pswd))
	if err != nil {
		return err
	}
	return nil
}
