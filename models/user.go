package models

import (
	"GYM-App/db"
	"errors"
)

type User struct {
	ID       int64
	UserName string
	Password string `binding:"required"`
}

// THIS IS NOT SAFE AT ALL. ITS JUST A TEST APP WITH A SIMPLE PASS CHECK (NEVER USER REAL PASS OR REAL DATA IN THE APP)
func (u *User) ValidateCredentials() error {
	query := "SELECT id, username FROM users WHERE password = ?"
	err := db.DB.QueryRow(query, u.Password).Scan(&u.ID, &u.UserName)
	if err != nil {
		return errors.New("credentials invalid")
	}

	return nil

}

// THIS IS NOT SAFE AT ALL. ITS JUST A TEST APP WITH A SIMPLE PASS CHECK (NEVER USER REAL PASS OR REAL DATA IN THE APP)
