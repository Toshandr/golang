package order

type Order struct {
	id     int
	userID int
	amount float64
	status string
}

func CreateOrder(ID int, UserID int, Amount float64, Status string) *Order {
	tmp := Order{ID, UserID, Amount, Status} //проверка позже
	return &tmp
}

func (o *Order) GetID() int {
	return o.id
}
func (o *Order) GetUserID() int {
	return o.userID
}
func (o *Order) GetAmount() float64 {
	return o.amount
}
func (o *Order) GetStatus() string {
	return o.status
}

func (o *Order) SetStatus(newStatus string) {
	o.status = newStatus
}
