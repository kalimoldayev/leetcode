package main

import "fmt"

func main()  {
	symbol := "LVIII"
	fmt.Println(romanInteger(symbol))

    //LVIII - 58
    //IIII - 4
    // MCMXCIV - 1994
}

func romanInteger(s string) int  {
	data := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

	result := 0

	for i := 0; i < len(s); i++ {
        prev := 0
        current := data[string(s[i])]
        next := 0

        if i > 0 {
            prev = data[string(s[i-1])]
        }
        if i != len(s) - 1 {
            if val, ok := data[string(s[i + 1])]; ok { // проходимся до последнего
                next = val
            }
        }

        if current >= next   { // если больше или равен
            if prev < current {
                current -= prev
            }
            result += current
        }
	}

	return result
}