package ordermanager

import (
	"manager/mainErrors"
	. "manager/order"
)

type OrderManager struct {
	orders []Order
	nextID int //id след товара
}

func CreateManager() *OrderManager {
	a := []Order{}
	return &OrderManager{a, 1}
}

func CreateManagerFewOrders(orders ...Order) *OrderManager {
	return &OrderManager{orders, len(orders) + 1}
}

func CreateManagerSliceOrders(orders []Order) *OrderManager {
	return &OrderManager{orders, len(orders) + 1}
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

func (m *OrderManager) UpdateStatus(newStatus string, id int) (ok bool, err error) {
	if id <= 0 {
		return false, mainErrors.ErrorOrderIDLessZero
	}
	indOrder := -1
	for i, v := range m.orders {
		if v.GetID() == id {
			indOrder = i
		}
	}
	if indOrder == -1 {
		return false, mainErrors.ErrorOrderID
	}
	switch newStatus {
	case "paid":
		if m.orders[indOrder].GetStatus() == "paid" {
			return false, mainErrors.ErrorOldStatus
		}
		m.orders[indOrder].SetStatus("paid")
		return true, nil
	case "new":
		if m.orders[indOrder].GetStatus() == "new" {
			return false, mainErrors.ErrorOldStatus
		}
		m.orders[indOrder].SetStatus("new")
		return true, nil
	case "cancelled":
		if m.orders[indOrder].GetStatus() == "cancelled" {
			return false, mainErrors.ErrorOldStatus
		}
		m.orders[indOrder].SetStatus("cancelled")
		return true, nil

	default:
		return false, mainErrors.ErrorUnknownStatus //изменить эту строку когда вернусь тут должен быть блок default и новая ошибка
	}
}
