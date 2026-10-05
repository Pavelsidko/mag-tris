package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int

	fmt.Print("Введите количество студентов: ")
	fmt.Scan(&n)

	count3 := 0
	count4 := 0
	count5 := 0

	min := 100
	max := 0

	fmt.Println("Результаты студентов:")

	for i := 0; i < n; i++ {
		pullUps := rand.Intn(21)

		fmt.Printf("Студент %d: %d подтягиваний\n", i+1, pullUps)

		if pullUps >= 12 && pullUps < 14 {
			count3++
		} else if pullUps >= 14 && pullUps < 16 {
			count4++
		} else if pullUps >= 16 {
			count5++
		}

		if pullUps < min {
			min = pullUps
		}

		if pullUps > max {
			max = pullUps
		}
	}

	fmt.Printf("\nСдали на 3: %d студентов\n", count3)
	fmt.Printf("Сдали на 4: %d студентов\n", count4)
	fmt.Printf("Сдали на 5: %d студентов\n", count5)
	fmt.Printf("Минимальное количество подтягиваний: %d\n", min)
	fmt.Printf("Максимальное количество подтягиваний: %d\n", max)
}
