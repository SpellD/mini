package main

import (
	"fmt"
)

func Mounth(a int) {
	switch a {
	case 1, 2, 3:
		fmt.Println("Зима\n")
	case 4, 5, 6:
		fmt.Println("Весна\n")
	case 7, 8, 9:
		fmt.Println("Лето\n")
	case 10, 11, 12:
		fmt.Println("Осень\n")
	default:
		fmt.Println("Такого месяца не существует\n")
	}
}

func Year(year int) {

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {

		fmt.Println("Год является високосным годом!")

	} else {

		fmt.Println("Год не является високосным годом!")

	}
}

func Convert(rub float64) {
	var usd float64 = 75
	fmt.Printf("У вас %.2f$", rub/usd)
}

func MaxArr(arr []int) int {
	max := arr[0]
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}
	return max
}

func SimpleCalc(a int, b float64, c float64) float64 {

	var res float64

	switch a {
	default:
		fmt.Printf("Неизвестная команда - %d", a)
	case 1:
		res = b + c
	case 2:
		res = b - c
	case 3:
		res = b * c
	case 4:
		res = b / c
	}
	return res
}

func main() {

}
