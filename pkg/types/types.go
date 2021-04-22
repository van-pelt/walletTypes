package types

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
	ID       string
	Amount   Money
	Category PaymentCategory
	Status   PaymentStatus
}
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
