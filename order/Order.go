package order

import "manager/mainErrors"

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

func (o *Order) UpdateStatus(newStatus string) (ok bool, err error) {
	switch newStatus {
	case "paid":
		if o.status == "paid" {
			return false, mainErrors.ErrorOldStatus
		}
		o.status = "paid"
		return true, nil
	case "new":
		if o.status == "new" {
			return false, mainErrors.ErrorOldStatus
		}
		o.status = "new"
		return true, nil
	case "cancelled":
		if o.status == "cancelled" {
			return false, mainErrors.ErrorOldStatus
		}
		o.status = "cancelled"
		return true, nil
	}
	return false, nil //изменить эту строку когда вернусь тут должен быть блок default и новая ошибка
}
