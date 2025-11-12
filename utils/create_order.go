package utils

import (
	"fmt"
)

func NewOrder() {
	var newOrder Order
	fmt.Println("Введите номер заказа:")
	fmt.Scan(&newOrder.OrderID)

	fmt.Println("Введите номер пользователя:")
	fmt.Scan(&newOrder.UserID)

	fmt.Println("Введите количество дней хранения:")
	fmt.Scan(&newOrder.StorageDays)

	add_order_to_json(newOrder)

}
