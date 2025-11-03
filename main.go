package main

import (
	"Route/utils"
	"fmt"
	"os"
	"slices"
)

type Order struct {
	order_id     uint8 `json:"order_id"`
	user_id      uint8 `json:"user_id"`
	storage_days uint8 `json:"storage_days"`
}

const (
	base_console_print = `Список доступных команд:
	1. Принять заказ от курьера
	2. Вернуть заказ курьеру 
	3. Выдать заказы и принять возвраты клиента
	4. Получить список заказов
	5. Получить список возвратов
	6. Получить историю заказов
	exit для выхода
	`
	datafile_name = "orders.json"
)

func get_user_input() string {
	var user_input string
	fmt.Println(base_console_print)
	fmt.Fscan(os.Stdin, &user_input)
	return user_input
}

func error_input() {
	fmt.Println("Ошибочный ввод. Введен недопустимый номер")
}

func main() {
	available_inputs := []string{"1", "2", "3", "4", "5", "6", "Принять заказ от курьера", "Вернуть заказ курьеру"}
	input_symbol := ""
	for input_symbol != "exit" {
		input_symbol = get_user_input()
		if slices.Contains(available_inputs, input_symbol) {
			switch input_symbol {
			case "1":
				utils.NewOrder()
			case "2":
				return
			case "3":
				return
			case "4":
				return
			case "5":
				return
			case "6":
				return
			}
		} else {
			error_input()
		}
	}
}
