package utils

import (
	"fmt"
	"os"
)

func GetordersList() {

	var clientID uint64
	fmt.Println("Введите номер клиента: ")
	fmt.Fscan(os.Stdin, &clientID)

	all_orders, err := get_data_from_json()
	if err != nil {
		fmt.Errorf("Ошибка при чтении данных: %v", err)
	}

	// Выводим данные по фильтру
	fmt.Println("Список заказов клиента: ")
	for _, order := range all_orders {
		if order.UserID == clientID {
			fmt.Printf("OrderID: %d, UserID: %d, StorageDate: %s\n", order.OrderID, order.UserID, order.StorageDate)
		}
	}
	fmt.Println()
}
