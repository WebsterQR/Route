package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func add_order_to_json(newOrder Order) {
	fmt.Println("add_order_to_json with params:", newOrder)

	// Читаем то что уже лежит в файле
	currentOrders, err := get_data_from_json()
	if err != nil {
		fmt.Println("Ошибка при открытии / чтении файла")
	}

	// Добавляем новый заказ в список
	currentOrders = append(currentOrders, newOrder)

	// Записываем в файл обновленные данные
	write_data_to_json(currentOrders)
}

func get_data_from_json() ([]Order, error) {
	// Список в который сложим все что есть в файле
	var currentOrders []Order

	file, err := os.Open(datafile_name)
	if err != nil {
		fmt.Println("Ошибка при открытии файла")
		return currentOrders, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&currentOrders)
	if err != nil {
		fmt.Println("Ошибка при чтении данных | Ошибка при декодировании JSON")
		return currentOrders, err
	}

	return currentOrders, nil

}

func write_data_to_json(orders_data []Order) error {
	// Открываем файл на запись
	file, err := os.Create(datafile_name)
	if err != nil {
		return fmt.Errorf("Ошибка создания файла: %v", err)
	}
	defer file.Close()

	// Создаем encoder для потоковой записи
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	// Записываем данные
	err = encoder.Encode(orders_data)
	if err != nil {
		return fmt.Errorf("ошибка потоковой записи JSON: %v", err)
	}

	return nil
}
