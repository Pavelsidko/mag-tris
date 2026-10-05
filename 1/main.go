package main

import (
	"fmt"
	"math"
)

func main() {
	var length, width, height float64
	var repair string

	fmt.Print("Введите длину комнаты: ")
	fmt.Scan(&length)

	fmt.Print("Введите ширину комнаты: ")
	fmt.Scan(&width)

	fmt.Print("Введите высоту потолка: ")
	fmt.Scan(&height)

	fmt.Print("Вид ремонта (обои/краска): ")
	fmt.Scan(&repair)

	perimeter := 2 * (length + width)
	wallArea := perimeter * height

	if repair == "обои" {
		var rollLength, rollWidth, rollPrice float64

		fmt.Print("Введите длину рулона обоев: ")
		fmt.Scan(&rollLength)

		fmt.Print("Введите ширину рулона обоев: ")
		fmt.Scan(&rollWidth)

		fmt.Print("Введите стоимость рулона: ")
		fmt.Scan(&rollPrice)

		rollArea := rollLength * rollWidth
		rolls := math.Ceil(wallArea / rollArea)
		totalCost := rolls * rollPrice

		fmt.Printf("Необходимо рулонов: %.0f\n", rolls)
		fmt.Printf("Стоимость: %.2f руб.\n", totalCost)

	} else if repair == "краска" {
		var paintConsumption, paintPrice float64

		fmt.Print("Введите расход краски на 1 м2: ")
		fmt.Scan(&paintConsumption)

		fmt.Print("Введите стоимость 1 литра краски: ")
		fmt.Scan(&paintPrice)

		paintLiters := wallArea * paintConsumption
		totalCost := paintLiters * paintPrice

		fmt.Printf("Необходимо краски: %.2f л\n", paintLiters)
		fmt.Printf("Стоимость: %.2f руб.\n", totalCost)

	} else {
		fmt.Println("Неизвестный вид ремонта")
	}
}