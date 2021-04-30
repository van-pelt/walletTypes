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

func (a *Account) ToDump() string {
	return fmt.Sprintf("%d;%s;%d\n", a.ID, a.Phone, a.Balance)
}

func (f *Favorite) ToDump() string {
	return fmt.Sprintf("%s;%d;%s;%d;%s\n", f.ID, f.AccountID, f.Name, f.Amount, f.Category)
}

func (p *Payment) ToDump() string {
	return fmt.Sprintf("%s;%d;%d;%s;%s\n", p.ID, p.AccountID, p.Amount, p.Category, p.Status)
}
