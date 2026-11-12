package mainErrors

import "errors"

var (
	ErrorOrderExists     error = errors.New("Этот товар уже существует")
	ErrorOrderNotExists  error = errors.New("Такого товара не существует, проверьте вводимые данные")
	ErrorAmount          error = errors.New("Неккоректная цена заказа, цена должна быть положительной величиной")
	ErrorUserID          error = errors.New("Такого пользователя не существует, проверьте данные")
	ErrorOrderID         error = errors.New("Такого ID не существует, проверьте данные")
	ErrorOrderIDLessZero error = errors.New("ID - положительное число, проверьте данные")
	ErrorOldStatus       error = errors.New("Это прежний статус")
	ErrorUnknownStatus   error = errors.New("Неизвестный статус, проверьте данные")
)
