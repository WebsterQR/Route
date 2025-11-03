package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

const datafile_name = "orders_data.json"

type Order struct {
	OrderID     uint8 `json:"order_id"`
	UserID      uint8 `json:"user_id"`
	StorageDays uint8 `json:"storage_days"`
}

func NewOrder() {
	var newOrder Order
	fmt.Println("Введите номер заказа:")
	fmt.Scan(&newOrder.OrderID)

	fmt.Println("Введите номер пользователя:")
	fmt.Scan(&newOrder.UserID)

	fmt.Println("Введите количество дней хранения:")
	fmt.Scan(&newOrder.StorageDays)

	// Проверка на существование файла
	file, err := os.Open(datafile_name)
	var currentOrders []Order
	if err == nil {
		defer file.Close()

		// Считываем уже имеющиеся данные
		err = json.NewDecoder(file).Decode(&currentOrders)
		if err != nil {
			currentOrders = []Order{}
		}
	}
	currentOrders = append(currentOrders, newOrder)
	file, err = os.Create(datafile_name)
	if err != nil {
		fmt.Println("Ошибка при создании файла")
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(currentOrders)
	if err != nil {
		fmt.Println("Ошибка при записи")
		return
	} else {
		fmt.Println("Успешно записано")
	}

}

func RefundOrder() {
	fmt.Println("Вызвано 2")
}

func PushPrders() {
	fmt.Println("3")
}

func GetordersList() {
	fmt.Println("4")
}

func GetRefundsList() {
	fmt.Println("5")
}

func GetOrdersHistory() {
	fmt.Println("6")
}
