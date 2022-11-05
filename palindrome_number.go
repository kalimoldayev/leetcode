package main

import "fmt"

func main() {
	x := 121
	fmt.Println(palindromeNumber(x))
}

func palindromeNumber(number int) bool {

	turnNumber := 0
	copyNumber := number

	for number > 0 {
		digit := number % 10                   // остаток от деления, таким образом получаем последнюю цифру
		number = number / 10                   // убираем последнюю цифру
		turnNumber = turnNumber * 10  + digit  // меням порядок
 	}

	if copyNumber == turnNumber {
		return true
	}

	return false
}