package domain

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID        string    `json:"id" db:"id"`
	UserId    string    `json:"user_id" db:"user_id"`
	Balance   float64   `json:"balance" db:"balance"`
	Currency  string    `json:"currency" db:"currency"`
	CreatedAt time.Time `json:"created_at" db:"created_st"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewWallet(userid string) *Wallet {
	now := time.Now()

	return &Wallet{
		ID:        uuid.New().String(),
		UserId:    userid,
		Balance:   0,
		Currency:  "KES",
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (w *Wallet) Validate() error {

	if w.UserId == "" {
		return ErrInvalidRequest
	}

	if w.Balance < 0 {
		return ErrInvalidAmount
	}

	if w.Currency == "" {
		return ErrInvalidCurrency
	}
	return nil

}

func (w *Wallet) Credit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	w.Balance += amount
	w.UpdatedAt = time.Now()
	return nil

}

func (w *Wallet) Debit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if amount > w.Balance {
		return ErrInSufficientBalance
	}
	w.Balance -= amount
	w.UpdatedAt = time.Now()
	return nil
}
