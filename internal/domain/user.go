package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string     `json:"id" db:"id"`
	Email        string     `json:"email" db:"email"`
	FullName     string     `json:"full_name" db:"full_name"`
	PasswordHash string     `json:"_" db:"password_hash"`
	Phone        string     `json:"phone,omitempty" db:"phone"`
	IsActive     bool       `json:"is_active" db:"is_active"`
	IsVerified   bool       `json:"is_verified" db:"is_verified"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	LastLogin    *time.Time `json:"last_login,omitempty" db:"last_login"`
}

// factory function
func NewUser(email, fullname, passwordHash string) *User {
	now := time.Now()

	return &User{
		ID:           uuid.New().String(),
		Email:        email,
		FullName:     fullname,
		PasswordHash: passwordHash,
		IsActive:     true,
		IsVerified:   false,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// busines logic
func (u *User) Validate() error {

	if u.Email == "" {
		return ErrInvalidEmail
	}

	if !isValidEmail(u.Email) {
		return ErrInvalidEmail
	}

	if u.FullName == "" {
		return ErrInvalidName
	}

	if len(u.FullName) > 100 {
		return ErrInvalidName
	}
	return nil

}

func (u *User) IsActiveUser() bool {
	return u.IsActive
}

func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
}

func (u *User) Verify() {
	u.IsVerified = true
	u.UpdatedAt = time.Now()
}

func (u *User) UpdateLastLogin() {
	now := time.Now()

	u.LastLogin = &now
}

//helper function

func isValidEmail(email string) bool {
	return len(email) > 0 && len(email) < 255
}
