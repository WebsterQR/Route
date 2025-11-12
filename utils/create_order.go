package utils

import (
	"fmt"
	"os"
)

func NewOrder() {
	var order_id, user_id uint64
	var storage_days int
	fmt.Println("Введите номер заказа:")
	fmt.Fscan(os.Stdin, &order_id)

	fmt.Println("Введите номер пользователя:")
	fmt.Fscan(os.Stdin, &user_id)

	fmt.Println("Введите количество дней хранения:")
	fmt.Fscan(os.Stdin, &storage_days)

	var newOrder Order
	newOrder.OrderID = order_id
	newOrder.UserID = user_id
	newOrder.StorageDate = convert_storage_days_to_data(storage_days)
	add_order_to_json(newOrder)
	fmt.Println(newOrder)

}
