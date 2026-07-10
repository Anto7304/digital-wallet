package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)


func GenerateID()string{
	return uuid.New().String()
}

func GenerateReference(prefix string)string{
	timestamp := time.Now().UnixNano()
	random := uuid.New().String()[:8]
	return fmt.Sprintf("%s-%d-%s",prefix,timestamp,random)
}
func GenerateTransactionReference()string{
	return GenerateReference("TX")
}

func GenerateDepositReference()string{
	return GenerateReference("DEP")
}

func GenerateWithdrawalReference()string{
	return GenerateReference("WTH")
}