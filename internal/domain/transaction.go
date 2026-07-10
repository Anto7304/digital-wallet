package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeTransfer   TransactionType = "transfer"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
	ID          string                 `json:"id" db:"id"`
	Reference   string                 `json:"reference" db:"reference"`
	SenderID    *string                `json:"sender_id,omitempty" db:"sender_id"`
	ReceiverID  *string                `json:"receiver_id,omitempty" db:"receiver_id"`
	Amount      float64                `json:"amount" db:"amount"`
	Type        TransactionType        `json:"type" db:"type"`
	Category    string                 `json:"category" db:"category"`
	Description string                 `json:"description" db:"description"`
	Status      TransactionStatus      `json:"status" db:"status"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" db:"metadata"`
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	CompletedAt *time.Time             `json:"completed_at,omitempty" db:"completed_at"`
}

func NewTransaction(
	txType TransactionType,
	amount float64,
	reference string,
	category string,
	descryption string,

) *Transaction {
	if reference == "" {
		reference = GenerateReference(string(txType))
	}

	return &Transaction{
		ID:          uuid.New().String(),
		Reference:   reference,
		Amount:      amount,
		Type:        txType,
		Category:    category,
		Description: descryption,
		CreatedAt:   time.Now(),
	}
}

func (t *Transaction) Validate() error {
	if t.Reference == "" {
		return ErrInvalidRequest
	}

	if t.Amount <= 0 {
		return ErrInvalidAmount
	}

	if t.Type == "" {
		return ErrInvalidTransactiontype
	}

	if t.Type != TransactionTypeDeposit && t.Type != TransactionTypeTransfer && t.Type != TransactionTypeWithdrawal {
		return ErrInvalidTransactiontype
	}
	return nil

}

func (t *Transaction) Complete() {
	now := time.Now()
	t.Status = TransactionStatusCompleted
	t.CompletedAt = &now
}

func (t *Transaction) Fail() {
	now := time.Now()
	t.Status = TransactionStatusFailed
	t.CompletedAt = &now
}

func (t *Transaction) IsCompleted() bool {
	return t.Status == TransactionStatusCompleted
}

func (t *Transaction) IsPending() bool {
	return t.Status == TransactionStatusPending
}
