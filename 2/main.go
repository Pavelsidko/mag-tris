package main

import "fmt"

func main() {
	var speed int

	fmt.Print("Введите скорость автомобиля: ")
	fmt.Scan(&speed)

	excess := speed - 90

	if excess <= 0 {
		fmt.Println("Скорость автомобиля допустима на данном участке")
	} else if excess < 20 {
		fmt.Println("Превышение скорости менее 20 км/ч. Штраф не предусмотрен.")
	} else if excess < 40 {
		fmt.Println("Штраф: 500 рублей")
	} else if excess < 60 {
		fmt.Println("Штраф: 1500 рублей")
	} else if excess < 80 {
		fmt.Println("Штраф: 2500 рублей или лишение прав на 4 месяца")
	} else {
		fmt.Println("Штраф: 5000 рублей или лишение прав на полгода")
	}
}
