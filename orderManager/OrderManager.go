package ordermanager

import (
	"manager/mainErrors"
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
	if ID <= 0 {
		return o, mainErrors.ErrorOrderIDLessZero
	}
	for i := 0; i < len(m.orders); i++ {
		if m.orders[i].GetID() == ID {
			return m.orders[i], nil
		}
	}
	return o, mainErrors.ErrorOrderNotExists
}

func (m *OrderManager) DeleteOrder(ID int) (ok bool, err error) {
	if ID <= 0 {
		return false, mainErrors.ErrorOrderIDLessZero
	}

	flag := -1

	for i, v := range m.orders {
		if v.GetID() == ID {
			flag = i
		}
	}
	if flag == -1 {
		return false, mainErrors.ErrorAmount // ошибку новую сделаю потом времени нет
	}
	firstPart := m.orders[0:flag]
	secondPart := m.orders[flag+1:]

	m.orders = append(firstPart, secondPart...)

	return true, nil

}

func (m *OrderManager) GetUserOrders(UserID int) (ListOrders []Order, err error) {
	if UserID <= 0 {
		return ListOrders, mainErrors.ErrorUserID
	}
	for _, v := range m.orders {
		if v.GetUserID() == UserID {
			ListOrders = append(ListOrders, v)
		}
	}
	return ListOrders, nil
}
