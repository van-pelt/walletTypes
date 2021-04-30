package types

import (
	"fmt"
)

const (
	StatusOk        PaymentStatus = "OK"
	StatusFail      PaymentStatus = "FAIL"
	StatusInProgres PaymentStatus = "INPROGRESS"
)

type PaymentCategory string
type PaymentStatus string
type Money int64
type Phone string

type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
type Favorite struct {
	ID        string
	AccountID int64
	Name      string
	Amount    Money
	Category  PaymentCategory
}

func (a *Account) ToString() string {
	return fmt.Sprintf("%d;%s;%d|", a.ID, a.Phone, a.Balance)
}
