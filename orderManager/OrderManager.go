package ordermanager

import (
	"manager/mainErrors"
	_ "manager/mainErrors"
	. "manager/order"
)

type OrderManager struct {
	orders []Order
	nextID int //id след товара
}

func (m *OrderManager) addOrder(UserID int, Amount float64) (id int, err error) {

	if Amount <= 0 {
		return 0, mainErrors.ErrorAmount
	}

	if UserID <= 0 {
		return 0, mainErrors.ErrorUserID
	}

	tmp := CreateOrder(m.nextID, UserID, Amount, "new")
	m.orders = append(m.orders, *tmp)
	m.nextID++
	return tmp.GetID(), nil
}

func (m *OrderManager) GetOrder(ID int) (o Order, err error) {
	if ID < 0 {
		return o, mainErrors.ErrorOrderIDLessZero
	}
	for i := 0; i < len(m.orders); i++ {
		if m.orders[i].GetID() == ID {
			return m.orders[i], nil
		}
	}
	return o, mainErrors.ErrorOrderNotExists
}
