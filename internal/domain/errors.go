package domain

import "errors"

//user
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserInactive       = errors.New("user account is an active")
	ErrUserNotVerified    = errors.New("user account not verified")
)

//validation error

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidName  = errors.New("invalid name")
)

//waalet errors
var(
	ErrInvalidRequest = errors.New("Wallet not found")
	ErrInSufficientBalance = errors.New("insufficient balance")
	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidCurrency = errors.New("invalid currency")
	ErrWalletAlreadyExist = errors.New("wallet already exist for this user")

)
//transaction
var(
	ErrTransactionNotfound =errors.New("transaction not found")
	ErrDuplicateTransaction = errors.New("duplicate transaction")
	ErrInvalidTransactiontype = errors.New("invalid transaction type")
	ErrInvalidTransactionStatus = errors.New("invalid transaction status")
	ErrTrnsactionAlreadyProcessed = errors.New("transaction already processed")
	ErrTransactionFailed =errors.New("transaction failed")
)