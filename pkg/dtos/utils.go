package dtos

import (
	"golang.org/x/crypto/bcrypt"
)

// Method for createUser model to hash the password
func (newUser *CreateUser) HashPassword(password string, user *CreateUser) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
