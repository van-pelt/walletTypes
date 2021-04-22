package types

const (
	StatusOk        Status = "OK"
	StatusFail      Status = "FAIL"
	StatusInProgres Status = "INPROGRESS"
)

type Currency string
type Category string
type Money int
type Status string
type Card struct {
	ID          int
	PAN         string
	Balance     Money
	MinBalance  Money
	Currency    Currency
	Color       string
	Name        string
	Active      bool
	MyInterface interface{}
	Status      string
}

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
